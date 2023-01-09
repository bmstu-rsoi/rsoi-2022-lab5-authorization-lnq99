package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/model"
)

func NewTicketService() Service {
	service := Service{
		Info: ServiceInfo{
			Name:       "Ticket",
			IP:         TicketServiceIP,
			ApiVersion: ApiVersion,
			Path:       "tickets",
		},
		Endpoints: []Endpoint{
			{"GET", "", GetTickets},
			{"POST", "", ForwardToTicketService},
			{"GET", ":ticketUid", GetTicket},
			{"DELETE", ":ticketUid", ForwardToTicketService},
		},
	}
	return service
}

func ForwardToTicketService(c *fiber.Ctx) error {
	addr := TicketServiceIP + c.OriginalURL()
	return proxy.Forward(addr)(c)
}

func GetTickets(c *fiber.Ctx) error {
	url := TicketServiceIP + c.OriginalURL()
	header := map[string]string{UsernameHeader: c.GetReqHeaders()[UsernameHeader]}

	r, err := CallServiceWithCircuitBreaker(
		ticketCb, "GET", url, header, nil)

	return fiberProcessResponse[[]model.TicketResponse](c, r.status, r.body, err)
}

func GetTicket(c *fiber.Ctx) error {
	url := TicketServiceIP + c.OriginalURL()
	header := map[string]string{UsernameHeader: c.GetReqHeaders()[UsernameHeader]}

	r, err := CallServiceWithCircuitBreaker(
		ticketCb, "GET", url, header, nil)

	return fiberProcessResponse[model.TicketResponse](c, r.status, r.body, err)
}
