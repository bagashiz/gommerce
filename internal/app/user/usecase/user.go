package usecase

import (
	"context"

	cityDom "github.com/bagashiz/gommerce/internal/app/city/domain"
	provDom "github.com/bagashiz/gommerce/internal/app/province/domain"
	"github.com/bagashiz/gommerce/internal/app/user/domain"
)

// UserUsecase is a struct that implements UserUsecase interface.
type UserUsecase struct {
	userRepo domain.UserRepository
	cityRepo cityDom.CityRepository
	provRepo provDom.ProvinceRepository
}

// New creates a new CategoryUsecase instance.
func New(userRepo domain.UserRepository, cityRepo cityDom.CityRepository, provRepo provDom.ProvinceRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepo,
		cityRepo,
		provRepo,
	}
}

// GetByID returns the User with the specified ID.
func (uu *UserUsecase) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	user, err := uu.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	province, err := uu.provRepo.GetByID(user.ProvinceID)
	if err != nil {
		return nil, err
	}

	city, err := uu.cityRepo.GetByID(user.CityID)
	if err != nil {
		return nil, err
	}

	user.Province = province
	user.City = city

	return user, nil
}

// Update updates the User with the specified ID.
func (uu *UserUsecase) Update(ctx context.Context, user *domain.User) error {
	_, err := uu.GetByID(ctx, user.ID)
	if err != nil {
		return err
	}

	return uu.userRepo.Update(ctx, user)
}

// IsAdmin checks if the User with the specified ID is an admin.
func (uu *UserUsecase) IsAdmin(ctx context.Context, id uint) (bool, error) {
	return uu.userRepo.IsAdmin(ctx, id)
}
