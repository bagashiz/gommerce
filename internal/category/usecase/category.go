package usecase

import (
	"context"

	"github.com/bagashiz/gommerce/internal/category/domain"
)

// CategoryUsecase is a struct that implements CategoryUsecase interface.
type CategoryUsecase struct {
	repo domain.CategoryRepository
}

// New creates a new CategoryUsecase instance.
func New(repo domain.CategoryRepository) domain.CategoryUsecase {
	return &CategoryUsecase{
		repo,
	}
}

// Create stores a new Category.
func (cu *CategoryUsecase) Create(ctx context.Context, category *domain.Category) error {
	return cu.repo.Create(ctx, category)
}

// GetAll returns all Categories.
func (cu *CategoryUsecase) GetAll(ctx context.Context, page, limit int) ([]domain.Category, error) {
	return cu.repo.GetAll(ctx, page, limit)
}

// GetByID returns the Category with the specified ID.
func (cu *CategoryUsecase) GetByID(ctx context.Context, id uint) (*domain.Category, error) {
	return cu.repo.GetByID(ctx, id)
}

// Update updates the Category with the specified ID.
func (cu *CategoryUsecase) Update(ctx context.Context, category *domain.Category) error {
	_, err := cu.GetByID(ctx, category.ID)
	if err != nil {
		return err
	}

	return cu.repo.Update(ctx, category)
}

// Delete removes the Category with the specified ID.
func (cu *CategoryUsecase) Delete(ctx context.Context, id uint) error {
	_, err := cu.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return cu.repo.Delete(ctx, id)
}
