package exoplanets

import (
	"context"
	"exoplanetservice/models/dto"
)

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
