package domain

import (
	"context"
)

//go:generate mockgen -source=category.go -destination=mock/category.go -package=mock

// Category is a struct that represents the Category of products.
type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// CategoryRepository is an interface that provides access to the Category storage.
type CategoryRepository interface {
	// Create stores a new Category.
	Create(ctx context.Context, category *Category) error
	// GetAll returns all Categories.
	GetAll(ctx context.Context, page, limit int) ([]Category, error)
	// GetByID returns the Category with the specified ID.
	GetByID(ctx context.Context, id uint) (*Category, error)
	// Update updates the Category with the specified ID.
	Update(ctx context.Context, category *Category) error
	// Delete removes the Category with the specified ID.
	Delete(ctx context.Context, id uint) error
}

// CategoryUsecase is an interface that provides business logic for Category.
type CategoryUsecase interface {
	// Create stores a new Category.
	Create(ctx context.Context, category *Category) error
	// GetAll returns all Categories.
	GetAll(ctx context.Context, page, limit int) ([]Category, error)
	// GetByID returns the Category with the specified ID.
	GetByID(ctx context.Context, id uint) (*Category, error)
	// Update updates the Category with the specified ID.
	Update(ctx context.Context, category *Category) error
	// Delete removes the Category with the specified ID.
	Delete(ctx context.Context, id uint) error
}
