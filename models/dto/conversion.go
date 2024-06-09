package dto

import (
	"exoplanetservice/models/dao"
	"time"
)

func ExoplanetToDTO(exoplanet *dao.Exoplanets) *Exoplanet {
	return &Exoplanet{
		ID:                exoplanet.ID,
		Name:              exoplanet.Name,
		Description:       exoplanet.Description,
		DistanceFromEarth: exoplanet.DistanceFromEarth,
		Radius:            exoplanet.Radius,
		Mass:              exoplanet.Mass,
		TypeOfExoplanet:   exoplanet.TypeOfExoplanet,
		CreatedAt:         exoplanet.CreatedAt,
		UpdatedAt:         exoplanet.UpdatedAt,
	}
}

func (r *ExoplanetRequest) ToDaoObject() *dao.Exoplanets {
	return &dao.Exoplanets{
		ID:                r.ID,
		Name:              r.Name,
		Description:       r.Description,
		DistanceFromEarth: r.DistanceFromEarth,
		Radius:            r.Radius,
		Mass:              r.Mass,
		TypeOfExoplanet:   r.TypeOfExoplanet,
		CreatedAt:         time.Now().UnixMilli(),
		UpdatedAt:         time.Now().UnixMilli(),
	}
}

// ConvertToExoplanetResponseList converts a list of ExoplanetsDAO to a list of ExoplanetDTO.
func ConvertToExoplanetResponseList(exoplanets []*dao.Exoplanets) []*Exoplanet {
	var exoplanetResponses []*Exoplanet
	for _, exoplanet := range exoplanets {
		exoplanetResponses = append(exoplanetResponses, ExoplanetToDTO(exoplanet))
	}
	return exoplanetResponses
}
