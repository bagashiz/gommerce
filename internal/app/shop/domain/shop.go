package domain

import (
	"context"
)

// Shop is a struct that represents the User's shop.
type Shop struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
	UserID         uint   `json:"user_id"`
}

// ShopRepository is an interface that provides access to the Shop storage.
type ShopRepository interface {
	// GetAll returns all Shops.
	GetAll(ctx context.Context, page, limit int, name string) ([]Shop, error)
	// GetUserShop returns the Shop with the specified User ID.
	GetUserShop(ctx context.Context, userID uint) (*Shop, error)
	// GetByID returns the Shop with the specified ID.
	GetByID(ctx context.Context, id uint) (*Shop, error)
	// Update updates the Shop with the specified ID.
	Update(ctx context.Context, shop *Shop) error
}

// ShopUsecase is an interface that provides business logic for Shop.
type ShopUsecase interface {
	// GetAll returns all Categories.
	GetAll(ctx context.Context, page, limit int, name string) ([]Shop, error)
	// GetUserShop returns the Shop with the specified User ID.
	GetUserShop(ctx context.Context, userID uint) (*Shop, error)
	// GetByID returns the Shop with the specified ID.
	GetByID(ctx context.Context, id uint) (*Shop, error)
	// Update updates the Shop with the specified ID.
	Update(ctx context.Context, userID, id uint) error
}
