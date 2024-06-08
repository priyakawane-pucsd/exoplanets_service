package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"exoplanetservice/controller/handlers/ping"
	"exoplanetservice/controller/handlers/swagger"
	"exoplanetservice/logger"
	"exoplanetservice/service"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port               int
	GinMode            string
	ReadTimeoutMillis  int
	WriteTimeoutMillis int
	Ping               ping.Config
}

type Controller struct {
	server         *http.Server
	conf           *Config
	serviceFactory *service.ServiceFactory
}

func NewController(ctx context.Context, c *Config, sf *service.ServiceFactory) *Controller {
	return &Controller{conf: c, serviceFactory: sf}
}

func (c *Controller) registerHandlers(ctx context.Context, router gin.IRouter) {
	ping.NewHandler(ctx, &c.conf.Ping, c.serviceFactory.PingService).Register(router)
	swagger.NewHandler(ctx).Register(router)
}

func (c *Controller) Listen(ctx context.Context) error {
	router := gin.New()

	c.registerHandlers(ctx, router.Group("/exoplanetservice"))

	logger.Info("🌏 go-base-service started on 🌎 -> http://localhost:%d/", c.conf.Port)

	c.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", c.conf.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(c.conf.ReadTimeoutMillis) * time.Millisecond,
		WriteTimeout: time.Duration(c.conf.WriteTimeoutMillis) * time.Millisecond,
	}
	err := c.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Panic(ctx, "Failed to start server : %v", err)
		return err
	}
	return nil
}
