package repository

import (
	"context"

	"github.com/bagashiz/gommerce/internal/category/domain"
	"github.com/bagashiz/gommerce/internal/pkg/database"
	"github.com/bagashiz/gommerce/internal/pkg/database/dao"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
)

// CategoryRepository is a struct that implements CategoryRepository interface.
type CategoryRepository struct {
	conn database.DB
}

// New creates a new CategoryRepository instance.
func New(conn database.DB) domain.CategoryRepository {
	return &CategoryRepository{
		conn,
	}
}

// Create stores a new Category.
func (cr *CategoryRepository) Create(ctx context.Context, category *domain.Category) error {
	dao := cr.toDAO(category)

	result := cr.conn.DB().Create(&dao).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "duplicated key not allowed" {
			return helper.ErrDataAlreadyExists
		}

		return result.Error
	}

	return nil
}

// GetAll returns all Categories.
func (cr *CategoryRepository) GetAll(ctx context.Context, page, limit int) ([]domain.Category, error) {
	var (
		categories []domain.Category
		daos       []dao.Category
	)

	offset := (page - 1) * limit

	result := cr.conn.DB().Find(&daos).Limit(limit).Offset(offset).WithContext(ctx)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, dao := range daos {
		categories = append(categories, *cr.toDomain(&dao))
	}

	return categories, nil
}

// GetByID returns the Category with the specified ID.
func (cr *CategoryRepository) GetByID(ctx context.Context, id uint) (*domain.Category, error) {
	var dao dao.Category

	result := cr.conn.DB().First(&dao, id).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, helper.ErrDataNotFound
		}

		return nil, result.Error
	}

	category := cr.toDomain(&dao)

	return category, nil
}

// Update updates the Category with the specified ID.
func (cr *CategoryRepository) Update(ctx context.Context, category *domain.Category) error {
	dao := cr.toDAO(category)

	result := cr.conn.DB().Model(dao).Updates(&dao).Where("id = ?", dao.ID).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return helper.ErrDataNotFound
		}

		return result.Error
	}

	return nil
}

// Delete removes the Category with the specified ID.
func (cr *CategoryRepository) Delete(ctx context.Context, id uint) error {
	var dao dao.Category

	result := cr.conn.DB().Delete(&dao, id).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return helper.ErrDataNotFound
		}

		return result.Error
	}

	return nil
}

// toDomain converts a DAO Category to a Category.
func (cr *CategoryRepository) toDomain(category *dao.Category) *domain.Category {
	return &domain.Category{
		ID:   category.ID,
		Name: category.Name,
	}
}

// toDAO converts a Category to a DAO Category.
func (cr *CategoryRepository) toDAO(category *domain.Category) *dao.Category {
	return &dao.Category{
		Model: dao.Model{
			ID: category.ID,
		},
		Name: category.Name,
	}
}
