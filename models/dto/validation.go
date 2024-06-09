package dto

import (
	"exoplanetservice/utils"
)

const (
	GAS_GIANT   = "GasGiant"
	TERRESTRIAL = "Terrestrial"
)

var validExoplanetTypes = map[string]bool{
	GAS_GIANT:   true,
	TERRESTRIAL: true,
}

func (er *ExoplanetRequest) Validate() error {
	// Validate name
	if er.Name == "" {
		return utils.BAD_REQUEST_ERROR("name cannot be empty")
	}

	// Validate description
	if er.Description == "" {
		return utils.BAD_REQUEST_ERROR("description cannot be empty")
	}

	// Validate distance from Earth
	if er.DistanceFromEarth <= 10 || er.DistanceFromEarth >= 1000 {
		return utils.BAD_REQUEST_ERROR("invalid distance from earth, it should be between 10 and 1000 light years")
	}

	// Validate radius
	if er.Radius <= 0.1 || er.Radius >= 10 {
		return utils.BAD_REQUEST_ERROR("invalid radius, it should be between 0.1 and 10 Earth-radius units")

	}

	// Validate exoplanet type
	if !validExoplanetTypes[er.TypeOfExoplanet] {
		return utils.BAD_REQUEST_ERROR("invalid exoplanet type , should be %s/%s", GAS_GIANT, TERRESTRIAL)
	}

	// Validate mass (only if the type is Terrestrial)
	if er.TypeOfExoplanet == "Terrestrial" {
		if er.Mass <= 0.1 || er.Mass >= 10 {
			return utils.BAD_REQUEST_ERROR("invalid mass, it should be between 0.1 and 10 Earth-mass units for Terrestrial exoplanets")
		}
	} else if er.TypeOfExoplanet == "GasGiant" {
		if er.Mass != 0 {
			return utils.BAD_REQUEST_ERROR("mass should not be provided for GasGiant exoplanets")
		}
	}
	return nil
}
