package services

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/adjust/rmq/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/model"
)

var (
	RedisConnection  rmq.Connection
	BonusDeleteQueue rmq.Queue
)

const (
	prefetchLimit = 100
	pollDuration  = 200 * time.Millisecond
)

type DeleteBonusTask struct {
	Url string
	//Body   io.Reader
	//Header map[string]string
}

func NewBonusService() Service {
	s := Service{
		Info: ServiceInfo{
			Name:       "Bonus",
			IP:         BonusServiceIP,
			ApiVersion: ApiVersion,
			Path:       "privilege",
		},
		Endpoints: []Endpoint{
			{"GET", "", GetBonus},
			{"POST", "", PostBonus},
			{"DELETE", ":ticketUid", DeleteBonus},
		},
	}

	maxTries := 10
	var err error

	for i := 0; i < maxTries; i++ {
		RedisConnection, err = rmq.OpenConnection("redis", "tcp", "queue:6379", 1, nil)
		if err == nil {
			break
		}
	}

	for i := 0; i < maxTries; i++ {
		BonusDeleteQueue, err = RedisConnection.OpenQueue("bonus")
		if err == nil {
			break
		}
	}

	log.Println("queue", BonusDeleteQueue)

	BonusDeleteQueue.StartConsuming(prefetchLimit, pollDuration)

	BonusDeleteQueue.AddConsumerFunc("post bonus", consumeDeleteBonusTask)

	BonusDeleteQueue.SetPushQueue(BonusDeleteQueue)

	return s
}

func ForwardToBonusService(c *fiber.Ctx) error {
	addr := BonusServiceIP + c.OriginalURL()
	return proxy.Forward(addr)(c)
}

func GetBonus(c *fiber.Ctx) error {
	url := BonusServiceIP + c.OriginalURL()
	header := map[string]string{UsernameHeader: c.GetReqHeaders()[UsernameHeader]}

	r, err := CallServiceWithCircuitBreaker(
		bonusCb, "GET", url, header, nil)

	return fiberProcessResponse[model.PrivilegeInfoResponse](c, r.status, r.body, err)
}

func PostBonus(c *fiber.Ctx) error {
	return ForwardToBonusService(c)
}

func DeleteBonus(c *fiber.Ctx) error {
	url := BonusServiceIP + c.OriginalURL()
	//body := bytes.NewReader(c.Body())

	r, _ := CallServiceWithCircuitBreaker(
		bonusCb, "DELETE", url, nil, nil)

	if r.status == http.StatusServiceUnavailable {
		taskBytes, _ := json.Marshal(DeleteBonusTask{
			Url: url,
		})
		log.Println(url)
		BonusDeleteQueue.PublishBytes(taskBytes)
	}

	return c.SendStatus(http.StatusOK)
}

func consumeDeleteBonusTask(delivery rmq.Delivery) {
	var task DeleteBonusTask
	var err error

	err = json.Unmarshal([]byte(delivery.Payload()), &task)

	if err == nil {
		r, err := CallServiceWithCircuitBreaker(
			bonusCb, "DELETE", task.Url, nil, nil)
		
		if err == nil && r.status == http.StatusNoContent {
			delivery.Ack()
		} else {
			delivery.Push()
		}
	}
}
