package services

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	errors "github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/error"
	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/model"
)

func GetMe(c *fiber.Ctx) error {
	url1 := fmt.Sprintf("%s/%s/tickets", TicketServiceIP, ApiVersion)
	url2 := fmt.Sprintf("%s/%s/privilege", BonusServiceIP, ApiVersion)
	header := map[string]string{UsernameHeader: c.GetReqHeaders()[UsernameHeader]}

	var r model.UserInfoResponse

	r1, err := CallServiceWithCircuitBreaker(
		ticketCb, "GET", url1, header, nil)

	if err == nil {
		err = json.NewDecoder(r1.body).Decode(&r.Tickets)
		r1.body.Close()
	}

	if err != nil {
		return c.Status(r1.status).JSON(errors.ToErrResponse(err))
	}

	r2, err := CallServiceWithCircuitBreaker(
		bonusCb, "GET", url2, header, nil)

	if err == nil {
		err = json.NewDecoder(r2.body).Decode(&r.Privilege)
		r2.body.Close()
	}

	return c.JSON(r)
}
