package memory

import (
	"context"
	"exoplanetservice/models/dao"
	"exoplanetservice/models/filters"
	"exoplanetservice/utils"

	"github.com/google/uuid"
)

type Config struct {
}

type Repository struct {
	cfg        *Config
	exoplanets map[string]dao.Exoplanets
}

func NewRepository(ctx context.Context, cfg *Config) *Repository {
	return &Repository{
		cfg:        cfg,
		exoplanets: make(map[string]dao.Exoplanets),
	}
}

// GetExoplanets retrieves exoplanets with pagination.
func (r *Repository) GetExoplanets(ctx context.Context, filter *filters.ExoplanetFilter, limit int, offset int) ([]*dao.Exoplanets, error) {
	var result []*dao.Exoplanets
	i := 0
	for _, exoplanet := range r.exoplanets {
		if filter.Mass != 0 && exoplanet.Mass != filter.Mass {
			continue
		}
		if filter.Radius != 0 && exoplanet.Radius != filter.Radius {
			continue
		}
		if i >= offset && i < offset+limit {
			result = append(result, &exoplanet)
		}
		i++
	}
	return result, nil
}

// CreateExoplanets implements service.Repository.
func (r *Repository) CreateExoplanets(ctx context.Context, exoplanet *dao.Exoplanets) error {
	exoplanet.ID = uuid.NewString()
	r.exoplanets[exoplanet.ID] = *exoplanet
	return nil
}

// DeleteExoplanetById implements service.Repository.
func (r *Repository) DeleteExoplanetById(ctx context.Context, id string) error {
	if _, exists := r.exoplanets[id]; exists {
		delete(r.exoplanets, id)
		return nil
	}
	return utils.BAD_REQUEST_ERROR("invalid id `%s`", id)
}

// GetExoplanetById implements service.Repository.
func (r *Repository) GetExoplanetById(ctx context.Context, id string) (*dao.Exoplanets, error) {
	if planet, exists := r.exoplanets[id]; exists {
		return &planet, nil
	}
	return nil, utils.ERROR_DATABASE_RECORD_NOT_FOUND
}

// Ping implements service.Repository.
func (*Repository) Ping(ctx context.Context) error {
	return nil
}

// UpdateExoplanetById implements service.Repository.
func (r *Repository) UpdateExoplanetById(ctx context.Context, exoplanet *dao.Exoplanets, exoplanetId string) error {
	exoplanet.ID = exoplanetId
	r.exoplanets[exoplanet.ID] = *exoplanet
	return nil
}
