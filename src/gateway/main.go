package main

import (
	"gateway/config"
	"gateway/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello, World 👋!")
	//})

	apiVersion := "api/v1"

	{
		services.ApiVersion = apiVersion
		services.BonusServiceIP = cfg.Service.BonusUrl
		services.FlightServiceIP = cfg.Service.FlightUrl
		services.TicketServiceIP = cfg.Service.TicketUrl
	}

	//url := func(url, path string) string { return fmt.Sprintf("%s/%s/%s", url, apiVersion, path) }

	app.Get("manage/health", func(ctx *fiber.Ctx) error {
		ctx.Status(fiber.StatusOK)
		return nil
	})

	v1 := app.Group(apiVersion)
	{
		//v1.Get("flights", proxy.Forward(url(cfg.Service.FlightUrl, "flights")))
		//v1.All("tickets", proxy.Forward(url(cfg.Service.TicketUrl, "tickets")))
		//v1.All("privilege", proxy.Forward(url(cfg.Service.BonusUrl, "privilege")))
		//v1.All("tickets/:ticketUid", proxy.Forward(cfg.Service.TicketUrl+"/api/v1/tickets/:ticketUid"))
		v1.Get("me", services.GetMe)
	}

	s := services.FiberServer{app}
	{
		s.RegisterService(services.NewFlightService())
		s.RegisterService(services.NewTicketService())
		s.RegisterService(services.NewBonusService())
	}

	app.Listen(cfg.Server.Host + ":" + cfg.Server.Port)
}
