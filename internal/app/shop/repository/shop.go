package repository

import (
	"context"

	"github.com/bagashiz/gommerce/internal/app/shop/domain"
	"github.com/bagashiz/gommerce/internal/pkg/database"
	"github.com/bagashiz/gommerce/internal/pkg/database/dao"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
)

// ShopRepository is a struct that implements ShopRepository interface.
type ShopRepository struct {
	conn database.DB
}

// New creates a new ShopRepository instance.
func New(conn database.DB) domain.ShopRepository {
	return &ShopRepository{
		conn,
	}
}

// GetAll returns all Shops.
func (sr *ShopRepository) GetAll(ctx context.Context, page, limit int, name string) ([]domain.Shop, error) {
	var (
		shops []domain.Shop
		daos  []dao.Shop
	)

	offset := (page - 1) * limit

	result := sr.conn.DB().Find(&daos).Where("name LIKE ?", "%"+name+"%").Limit(limit).Offset(offset).WithContext(ctx)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, dao := range daos {
		shops = append(shops, *sr.toDomain(&dao))
	}

	return shops, nil
}

// GetUserShop returns the Shop with the specified User ID.
func (sr *ShopRepository) GetUserShop(ctx context.Context, userID uint) (*domain.Shop, error) {
	var dao dao.Shop

	result := sr.conn.DB().First(&dao).Where("user_id = ?", userID).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, helper.ErrDataNotFound
		}

		return nil, result.Error
	}

	shop := sr.toDomain(&dao)

	return shop, nil
}

// GetByID returns the Shop with the specified ID.
func (sr *ShopRepository) GetByID(ctx context.Context, id uint) (*domain.Shop, error) {
	var dao dao.Shop

	result := sr.conn.DB().First(&dao, id).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, helper.ErrDataNotFound
		}

		return nil, result.Error
	}

	shop := sr.toDomain(&dao)

	return shop, nil
}

// Update updates the Shop with the specified ID.
func (sr *ShopRepository) Update(ctx context.Context, shop *domain.Shop) error {
	dao := sr.toDAO(shop)

	result := sr.conn.DB().Model(dao).Updates(&dao).Where("id = ?", dao.ID).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return helper.ErrDataNotFound
		}

		return result.Error
	}

	return nil
}

// toDomain converts a DAO Shop to a Shop.
func (sr *ShopRepository) toDomain(shop *dao.Shop) *domain.Shop {
	return &domain.Shop{
		ID:             shop.ID,
		Name:           shop.Name,
		ProfilePicture: shop.ProfilePicture,
		UserID:         shop.UserID,
	}
}

// toDAO converts a Shop to a DAO Shop.
func (sr *ShopRepository) toDAO(shop *domain.Shop) *dao.Shop {
	return &dao.Shop{
		Model: dao.Model{
			ID: shop.ID,
		},
		Name:           shop.Name,
		ProfilePicture: shop.ProfilePicture,
		UserID:         shop.UserID,
	}
}
