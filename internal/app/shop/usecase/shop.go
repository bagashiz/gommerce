package usecase

import (
	"context"

	"github.com/bagashiz/gommerce/internal/app/shop/domain"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
)

// ShopUsecase is a struct that implements ShopUsecase interface.
type ShopUsecase struct {
	repo domain.ShopRepository
}

// New creates a new ShopUsecase instance.
func New(repo domain.ShopRepository) domain.ShopUsecase {
	return &ShopUsecase{
		repo,
	}
}

// GetAll returns all Shops.
func (su *ShopUsecase) GetAll(ctx context.Context, page, limit int, name string) ([]domain.Shop, error) {
	return su.repo.GetAll(ctx, page, limit, name)
}

// GetUserShop returns the Shop with the specified User ID.
func (su *ShopUsecase) GetUserShop(ctx context.Context, userID uint) (*domain.Shop, error) {
	return su.repo.GetUserShop(ctx, userID)
}

// GetByID returns the Shop with the specified ID.
func (su *ShopUsecase) GetByID(ctx context.Context, id uint) (*domain.Shop, error) {
	return su.repo.GetByID(ctx, id)
}

// Update updates the Shop with the specified ID.
func (su *ShopUsecase) Update(ctx context.Context, userID, id uint) error {
	shop, err := su.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if shop.UserID != userID {
		return helper.ErrForbidden
	}

	return su.repo.Update(ctx, shop)
}
