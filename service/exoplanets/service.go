package exoplanets

import (
	"context"
	"exoplanetservice/models/dao"
)

type Config struct {
}

type Repository interface {
	CreateExoplanets(ctx context.Context, exoplanet *dao.Exoplanets) error
	GetExoplanets(ctx context.Context, limit, offset int) ([]*dao.Exoplanets, error)
	GetExoplanetById(ctx context.Context, id string) (*dao.Exoplanets, error)
	UpdateExoplanetById(ctx context.Context, exoplanet *dao.Exoplanets, exoplanetId string) error
	DeleteExoplanetById(ctx context.Context, id string) error
}

type Service struct {
	config *Config
	repo   Repository
}

func NewService(ctx context.Context, cfg *Config, repo Repository) *Service {
	return &Service{config: cfg, repo: repo}
}
