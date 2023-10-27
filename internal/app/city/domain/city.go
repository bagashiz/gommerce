package domain

// City is a struct that represents the User's address city.
type City struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

// CityRepository is an interface that provides access to the City storage.
type CityRepository interface {
	// GetAll returns all Cities.
	GetAll(provinceID string) ([]City, error)
	// GetByID returns the City with the specified ID.
	GetByID(id string) (*City, error)
}

// CityUsecase is an interface that provides business logic for City.
type CityUsecase interface {
	// GetAll returns all Cities.
	GetAll(provinceID string) ([]City, error)
	// GetByID returns the City with the specified ID.
	GetByID(provinceID, cityID string) (*City, error)
}
