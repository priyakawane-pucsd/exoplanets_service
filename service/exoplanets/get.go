package exoplanets

import (
	"context"
	"exoplanetservice/models/dto"
	"exoplanetservice/models/filters"
	"math"
)

// GetExoplanets retrieves a paginated list of exoplanets.
func (s *Service) GetExoplanets(ctx context.Context, filter *filters.ExoplanetFilter, limit, offset int) (*dto.ListExoplanetResponse, error) {
	exoplanets, err := s.repo.GetExoplanets(ctx, filter, limit, offset)
	if err != nil {
		return nil, err
	}
	response := &dto.ListExoplanetResponse{
		Exoplanets: dto.ConvertToExoplanetResponseList(exoplanets),
	}
	return response, nil
}

func (s *Service) GetExoplanetById(ctx context.Context, exoplanetId string) (*dto.ExoplanetByIdResponse, error) {
	exoplanet, err := s.repo.GetExoplanetById(ctx, exoplanetId)
	if err != nil {
		return nil, err
	}
	response := &dto.ExoplanetByIdResponse{
		Exoplanet: *dto.ExoplanetToDTO(exoplanet),
	}
	return response, nil
}

func (s *Service) CalculateFuelEstimation(ctx context.Context, exoplanetId string, crewCapacity int) (*dto.FuelEstimationResponse, error) {
	exoplanet, err := s.repo.GetExoplanetById(ctx, exoplanetId)
	if err != nil {
		return nil, err
	}
	var gravity float64
	switch exoplanet.TypeOfExoplanet {
	case dto.GAS_GIANT:
		gravity = (0.5 / (math.Pow(exoplanet.Radius, 2)))
	case dto.TERRESTRIAL:
		gravity = (exoplanet.Mass / (math.Pow(exoplanet.Radius, 2)))
	}
	estimatedFuel := float64(exoplanet.DistanceFromEarth) / (math.Pow(gravity, 2)) * float64(crewCapacity)
	return &dto.FuelEstimationResponse{EstimatedFuel: estimatedFuel}, nil
}
