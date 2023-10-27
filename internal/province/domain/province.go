package domain

// Province is a struct that represents the User's address province.
type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ProvinceRepository is an interface that provides access to the Province storage.
type ProvinceRepository interface {
	// GetAll returns all Provinces.
	GetAll() ([]Province, error)
	// GetByID returns the Province with the specified ID.
	GetByID(id string) (*Province, error)
}

// ProvinceUsecase is an interface that provides business logic for Province.
type ProvinceUsecase interface {
	// GetAll returns all Provinces.
	GetAll() ([]Province, error)
	// GetByID returns the Province with the specified ID.
	GetByID(id string) (*Province, error)
}
