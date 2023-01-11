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

	{
		services.ApiVersion = "v1"
		services.BonusServiceIP = cfg.Service.BonusUrl
		services.FlightServiceIP = cfg.Service.FlightUrl
		services.TicketServiceIP = cfg.Service.TicketUrl
	}

	//url := func(url, path string) string { return fmt.Sprintf("%s/%s/%s", url, apiVersion, path) }

	app := fiber.New()

	app.Get("manage/health", func(ctx *fiber.Ctx) error {
		ctx.Status(fiber.StatusOK)
		return nil
	})
	//app.Get("api/v1/authorize", services.Authorize)
	//app.Get("api/v1/callback", services.Callback)

	api := app.Group("api", services.IsAuthenticated)

	v1 := api.Group("v1")
	{
		//v1.Get("flights", proxy.Forward(url(cfg.Service.FlightUrl, "flights")))
		//v1.All("tickets", proxy.Forward(url(cfg.Service.TicketUrl, "tickets")))
		//v1.All("privilege", proxy.Forward(url(cfg.Service.BonusUrl, "privilege")))
		//v1.All("tickets/:ticketUid", proxy.Forward(cfg.Service.TicketUrl+"/api/v1/tickets/:ticketUid"))
		v1.Get("me", services.GetMe)
	}

	s := services.FiberRouter{api}
	{
		s.RegisterService(services.NewFlightService())
		s.RegisterService(services.NewTicketService())
		s.RegisterService(services.NewBonusService())
	}

	app.Listen(cfg.Server.Host + ":" + cfg.Server.Port)
}
