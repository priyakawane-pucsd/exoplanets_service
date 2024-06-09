package repository

import (
	"context"
	"exoplanetservice/repository/memory"
	"exoplanetservice/repository/mongo"
	"exoplanetservice/service"
)

const (
	MONGO  = "mongo"
	MEMORY = "memory"
)

type Config struct {
	Name   string
	Mongo  mongo.Config
	Memory memory.Config
}

type Repository interface {
	service.Repository
}

func NewRepository(ctx context.Context, cfg *Config) service.Repository {
	switch cfg.Name {
	case MONGO:
		return mongo.NewRepository(ctx, &cfg.Mongo)
	case MEMORY:
		return memory.NewRepository(ctx, &cfg.Memory)
	}
	panic("invalid repository name , provided `" + cfg.Name + "` expected + `" + MONGO + "`")
}
