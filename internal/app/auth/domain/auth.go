package domain

import (
	"context"

	cityDom "github.com/bagashiz/gommerce/internal/app/city/domain"
	provinceDom "github.com/bagashiz/gommerce/internal/app/province/domain"
	userDom "github.com/bagashiz/gommerce/internal/app/user/domain"
)

// AuthRepository is an interface that provides access to the User storage.
type AuthRepository interface {
	// Create stores a new User.
	Create(ctx context.Context, user *userDom.User) error
	// GetByPhoneNumber returns the User with the specified phone number.
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*userDom.User, error)
}

// AuthUsecase is an interface that provides business logic for User.
type AuthUsecase interface {
	// Register stores a new User.
	Register(ctx context.Context, user *userDom.User) error
	// Login returns User data and access token.
	Login(ctx context.Context, phoneNumber, password string) (*userDom.User, *cityDom.City, *provinceDom.Province, string, error)
}
