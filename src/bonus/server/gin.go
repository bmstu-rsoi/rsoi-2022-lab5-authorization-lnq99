package server

import (
	"net/http"

	"bonus/config"
	"bonus/controller"
	"bonus/service"

	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/server"

	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	server.BaseServer
	engine   *gin.Engine
	handlers *controller.GinController
	config   *config.ServerConfig
}

func NewGinServer(service service.Service, cfg *config.ServerConfig) server.Server {
	// engine.SetMode(engine.ReleaseMode)
	engine := gin.New()

	// Middleware
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(limits.RequestSizeLimiter(1000))

	ctrl := controller.NewGinController(service)
	s := GinServer{
		engine:   engine,
		handlers: ctrl,
		config:   cfg,
	}
	s.SetupRouter()
	return &s
}

func (s *GinServer) SetupRouter() {
	r := s.engine
	v1 := r.Group("api/v1")
	p := v1.Group("privilege")
	{
		p.GET("", s.handlers.ListPrivilegeHistories)
		p.POST("", s.handlers.UpdateBalanceAndHistory)
		p.DELETE(":ticketUid", s.handlers.RevertBalanceAndHistory)
	}
	r.GET("/manage/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}

func (s *GinServer) Run() {
	cfg := server.DefaultConfig()
	cfg.Addr = s.config.Host + ":" + s.config.Port
	cfg.Handler = s.engine
	s.InitHttpServer(cfg)
	s.BaseServer.Run()
}
