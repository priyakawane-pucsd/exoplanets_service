package dao

type Exoplanets struct {
	ID                string  `bson:"_id,omitempty"`
	Name              string  `bson:"name"`
	Description       string  `bson:"description"`
	DistanceFromEarth float64 `bson:"distanceFromEarth"`
	Radius            float64 `bson:"radius"`
	Mass              float64 `bson:"mass,omitempty"`
	TypeOfExoplanet   string  `bson:"typeOfExoplanet"`
	CreatedAt         int64   `bson:"createdAt,omitempty"`
	UpdatedAt         int64   `bson:"updatedAt,omitempty"`
}
