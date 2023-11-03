package domain

import (
	"context"
)

// Address is a struct that represents the Address account.
type Address struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Receiver    string `json:"receiver"`
	PhoneNumber string `json:"phone_number"`
	Details     string `json:"details"`
	UserID      uint   `json:"user_id"`
}

// AddressRepository is an interface that provides access to the Address storage.
type AddressRepository interface {
	// Create stores a new Address.
	Create(ctx context.Context, addr *Address) error
	// GetAll returns all Addresses.
	GetAll(ctx context.Context, userID uint, title string) ([]Address, error)
	// GetByID returns the Address with the specified ID.
	GetByID(ctx context.Context, userID, id uint) (*Address, error)
	// Update updates the Address with the specified ID.
	Update(ctx context.Context, addr *Address) error
	// Delete removes the Address with the specified ID.
	Delete(ctx context.Context, userID, id uint) error
}

// AddressUsecase is an interface that provides business logic for Address.
type AddressUsecase interface {
	// Create stores a new Address.
	Create(ctx context.Context, addr *Address) error
	// GetAll returns all Addresses.
	GetAll(ctx context.Context, userID uint, title string) ([]Address, error)
	// GetByID returns the Address with the specified ID.
	GetByID(ctx context.Context, userID, id uint) (*Address, error)
	// Update updates the Address with the specified ID.
	Update(ctx context.Context, addr *Address) error
	// Delete removes the Address with the specified ID.
	Delete(ctx context.Context, userID, id uint) error
}
