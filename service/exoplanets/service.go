package exoplanets

import (
	"context"
	"exoplanetservice/models/dao"
	"exoplanetservice/models/dto"
	"math"
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

func (s *Service) CreateExoplanets(ctx context.Context, req *dto.ExoplanetRequest) (*dto.Exoplanet, error) {
	exoplanet := req.ToDaoObject()
	err := s.repo.CreateExoplanets(ctx, exoplanet)
	if err != nil {
		return nil, err
	}
	return dto.ExoplanetToDTO(exoplanet), nil
}

// GetExoplanets retrieves a paginated list of exoplanets.
func (s *Service) GetExoplanets(ctx context.Context, limit, offset int) (*dto.ListExoplanetResponse, error) {
	exoplanets, err := s.repo.GetExoplanets(ctx, limit, offset)
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

func (s *Service) UpdateExoplanetById(ctx context.Context, id string, updateRequest dto.ExoplanetRequest) error {
	exoplanet, err := s.repo.GetExoplanetById(ctx, id)
	if err != nil {
		return err
	}
	// Call repository to update exoplanet
	updateExoplanetReq := updateRequest.ToDaoObject()
	err = s.repo.UpdateExoplanetById(ctx, updateExoplanetReq, exoplanet.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteExoplanetById(ctx context.Context, examId string) error {
	err := s.repo.DeleteExoplanetById(ctx, examId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CalculateFuelEstimation(ctx context.Context, exoplanetId string, crewCapacity int) (float64, error) {
	exoplanet, err := s.repo.GetExoplanetById(ctx, exoplanetId)
	if err != nil {
		return 0, err
	}
	if exoplanet.TypeOfExoplanet == dto.GAS_GIANT {
		gravity := (0.5 / (math.Pow(exoplanet.Radius, 2)))
		return float64(exoplanet.DistanceFromEarth) / (math.Pow(gravity, 2)) * float64(crewCapacity), nil

	}
	gravity := (exoplanet.Mass / (math.Pow(exoplanet.Radius, 2)))
	return float64(exoplanet.DistanceFromEarth) / (math.Pow(gravity, 2)) * float64(crewCapacity), nil
}
