package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/model"
)

func NewFlightService() Service {
	service := Service{
		Info: ServiceInfo{
			Name:       "Flight",
			IP:         FlightServiceIP,
			ApiVersion: ApiVersion,
			Path:       "flights",
		},
		Endpoints: []Endpoint{
			{"GET", "", GetFlights},
			{"GET", ":flightNumber", GetFlight},
		},
	}
	return service
}

func ForwardToFlightService(c *fiber.Ctx) error {
	addr := FlightServiceIP + c.OriginalURL()
	return proxy.Forward(addr)(c)
}

func GetFlights(c *fiber.Ctx) error {
	url := FlightServiceIP + c.OriginalURL()

	r, err := CallServiceWithCircuitBreaker(
		flightCb, "GET", url, nil, nil)

	return fiberProcessResponse[model.PaginationResponse](c, r.status, r.body, err)
}

func GetFlight(c *fiber.Ctx) error {
	url := FlightServiceIP + c.OriginalURL()

	r, err := CallServiceWithCircuitBreaker(
		flightCb, "GET", url, nil, nil)

	return fiberProcessResponse[model.FlightResponse](c, r.status, r.body, err)
}
