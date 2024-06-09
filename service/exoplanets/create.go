package exoplanets

import (
	"context"
	"exoplanetservice/models/dto"
)

func (s *Service) CreateExoplanets(ctx context.Context, req *dto.ExoplanetRequest) (*dto.Exoplanet, error) {
	exoplanet := req.ToDaoObject()
	err := s.repo.CreateExoplanets(ctx, exoplanet)
	if err != nil {
		return nil, err
	}
	return dto.ExoplanetToDTO(exoplanet), nil
}
