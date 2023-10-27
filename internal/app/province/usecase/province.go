package usecase

import "github.com/bagashiz/gommerce/internal/app/province/domain"

// ProvinceUsecase is a struct that implements ProvinceUsecase interface.
type ProvinceUsecase struct {
	repo domain.ProvinceRepository
}

// New creates a new ProvinceUsecase instance.
func New(repo domain.ProvinceRepository) domain.ProvinceRepository {
	return &ProvinceUsecase{
		repo,
	}
}

// GetAll returns all Provinces.
func (pu *ProvinceUsecase) GetAll() ([]domain.Province, error) {
	return pu.repo.GetAll()
}

// GetByID returns the Province with the specified ID.
func (pu *ProvinceUsecase) GetByID(id string) (*domain.Province, error) {
	return pu.repo.GetByID(id)
}
