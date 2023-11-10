package usecase

import (
	"context"

	"github.com/bagashiz/gommerce/internal/app/auth/domain"
	cityDom "github.com/bagashiz/gommerce/internal/app/city/domain"
	provinceDom "github.com/bagashiz/gommerce/internal/app/province/domain"
	userDom "github.com/bagashiz/gommerce/internal/app/user/domain"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
	"github.com/bagashiz/gommerce/internal/pkg/token"
)

// AuthUsecase is a struct that implements AuthUsecase interface.
type AuthUsecase struct {
	authRepo   domain.AuthRepository
	userRepo   userDom.UserRepository
	cityRepo   cityDom.CityRepository
	provRepo   provinceDom.ProvinceRepository
	tokenMaker token.Token
}

// New creates a new AuthUsecase instance.
func New(authRepo domain.AuthRepository, userRepo userDom.UserRepository, cityRepo cityDom.CityRepository, provRepo provinceDom.ProvinceRepository, tokenMaker token.Token) domain.AuthUsecase {
	return &AuthUsecase{
		authRepo,
		userRepo,
		cityRepo,
		provRepo,
		tokenMaker,
	}
}

// Register stores a new User.
func (au *AuthUsecase) Register(ctx context.Context, user *userDom.User) error {
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return au.authRepo.Create(ctx, user)
}

// Login returns User data and access token.
func (au *AuthUsecase) Login(ctx context.Context, phoneNumber, password string) (*userDom.User, *cityDom.City, *provinceDom.Province, string, error) {
	user, err := au.authRepo.GetByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return nil, nil, nil, "", err
	}

	if err := helper.ComparePassword(password, user.Password); err != nil {
		return nil, nil, nil, "", err
	}

	city, err := au.cityRepo.GetByID(user.CityID)
	if err != nil {
		return nil, nil, nil, "", err
	}

	province, err := au.provRepo.GetByID(city.ProvinceID)
	if err != nil {
		return nil, nil, nil, "", err
	}

	token, err := au.tokenMaker.Create(user.ID, user.IsAdmin)
	if err != nil {
		return nil, nil, nil, "", err
	}

	return user, city, province, token, nil
}
