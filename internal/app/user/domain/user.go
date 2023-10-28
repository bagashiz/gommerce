package domain

import (
	"context"
	"time"

	cityDom "github.com/bagashiz/gommerce/internal/app/city/domain"
	provDom "github.com/bagashiz/gommerce/internal/app/province/domain"
)

// User is a struct that represents the User account.
type User struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Password    string            `json:"password"`
	PhoneNumber string            `json:"phone_number"`
	Email       string            `json:"email"`
	BirthDate   time.Time         `json:"birth_date" `
	About       string            `json:"about"`
	Job         string            `json:"job"`
	ProvinceID  string            `json:"province_id"`
	CityID      string            `json:"city_id"`
	Province    *provDom.Province `json:"province"`
	City        *cityDom.City     `json:"city"`
}

// UserRepository is an interface that provides access to the User storage.
type UserRepository interface {
	// GetByID returns the User with the specified ID.
	GetByID(ctx context.Context, id uint) (*User, error)
	// Update updates the User with the specified ID.
	Update(ctx context.Context, user *User) error
	// IsAdmin checks if the User with the specified ID is an admin.
	IsAdmin(ctx context.Context, id uint) (bool, error)
}

// UserUsecase is an interface that provides business logic for User.
type UserUsecase interface {
	// GetByID returns the User with the specified ID.
	GetByID(ctx context.Context, id uint) (*User, error)
	// Update updates the User with the specified ID.
	Update(ctx context.Context, user *User) error
}
