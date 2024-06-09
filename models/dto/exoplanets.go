package dto

type Exoplanet struct {
	ID                string  `json:"id,omitempty"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	DistanceFromEarth float64 `json:"distance_from_earth"`
	Radius            float64 `json:"radius"`
	Mass              float64 `json:"mass,omitempty"`
	TypeOfExoplanet   string  `json:"type_of_exoplanet"`
	CreatedAt         int64   `json:"created_at,omitempty"`
	UpdatedAt         int64   `json:"updated_at,omitempty"`
}

type ExoplanetRequest struct {
	ID                string  `json:"-"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	DistanceFromEarth float64 `json:"distance_from_earth"`
	Radius            float64 `json:"radius"`
	Mass              float64 `json:"mass,omitempty"`
	TypeOfExoplanet   string  `json:"type_of_exoplanet"`
}

type ExoplanetByIdResponse struct {
	Exoplanet Exoplanet `json:"exoplanet"`
}

// ListExoplanetResponse represents the response containing a list of exoplanets.
type ListExoplanetResponse struct {
	Exoplanets []*Exoplanet `json:"exoplanets"`
}

type FuelEstimationResponse struct {
	EstimatedFuel float64 `json:"estimatedFuel"`
}
