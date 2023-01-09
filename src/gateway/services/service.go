package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	errors "github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/error"
	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/model"
)

var (
	ApiVersion      = "api/v1"
	BonusServiceIP  = ""
	FlightServiceIP = ""
	TicketServiceIP = ""
	Client          = &http.Client{}
	UsernameHeader  = model.UsernameHeader
)

type Endpoint struct {
	Method  string
	Path    string
	Handler fiber.Handler
}

type ServiceInfo struct {
	Name       string
	IP         string
	ApiVersion string
	Path       string
}

type Service struct {
	Info      ServiceInfo
	Endpoints []Endpoint
}

type FiberServer struct {
	App *fiber.App
}

func (s FiberServer) RegisterRoute(e *Endpoint, prefix string) {
	s.App.Add(e.Method, prefix+e.Path, e.Handler)
}

func (s FiberServer) RegisterService(service Service) {
	prefix := service.Info.ApiVersion + "/" + service.Info.Path + "/"
	for _, e := range service.Endpoints {
		s.RegisterRoute(&e, prefix)
	}
}

func fiberProcessResponse[T any](c *fiber.Ctx, status int, body io.ReadCloser, err error) error {
	var r T

	if err == nil {
		err = json.NewDecoder(body).Decode(&r)
		body.Close()
	}

	c.Status(status)
	if err != nil {
		return c.JSON(errors.ToErrResponse(err))
	}
	return c.JSON(r)
}
