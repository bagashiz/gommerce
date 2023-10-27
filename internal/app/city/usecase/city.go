package usecase

import (
	"github.com/bagashiz/gommerce/internal/app/city/domain"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
)

// CityUsecase is a struct that implements CityUsecase interface.
type CityUsecase struct {
	repo domain.CityRepository
}

// New creates a new CityUsecase instance.
func New(repo domain.CityRepository) domain.CityUsecase {
	return &CityUsecase{
		repo,
	}
}

// GetAll returns all Cities.
func (pu *CityUsecase) GetAll(provinceID string) ([]domain.City, error) {
	return pu.repo.GetAll(provinceID)
}

// GetByID returns the City with the specified ID.
func (pu *CityUsecase) GetByID(provinceID, cityID string) (*domain.City, error) {
	if len(cityID) <= 2 {
		return nil, helper.ErrDataNotFound
	}

	if provinceID != cityID[:2] {
		return nil, helper.ErrDataNotFound
	}

	return pu.repo.GetByID(cityID)
}
