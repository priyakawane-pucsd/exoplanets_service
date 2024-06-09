package service

import (
	"context"
	"exoplanetservice/service/exoplanets"
	"exoplanetservice/service/ping"
)

type Config struct {
	Ping       ping.Config
	Exoplanets exoplanets.Config
}

type Repository interface {
	ping.Repository
	exoplanets.Repository
}

type ServiceFactory struct {
	PingService      *ping.Service
	ExoplanetService *exoplanets.Service
}

func NewServiceFactory(ctx context.Context, cfg *Config, repo Repository) *ServiceFactory {
	sf := ServiceFactory{}
	sf.PingService = ping.NewService(ctx, &cfg.Ping, repo)
	sf.ExoplanetService = exoplanets.NewService(ctx, &cfg.Exoplanets, repo)
	return &sf
}
