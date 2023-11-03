package repository

import (
	"context"

	"github.com/bagashiz/gommerce/internal/app/address/domain"
	"github.com/bagashiz/gommerce/internal/pkg/database"
	"github.com/bagashiz/gommerce/internal/pkg/database/dao"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
)

// AddressRepository is a struct that implements AddressRepository interface.
type AddressRepository struct {
	conn database.DB
}

// New creates a new AddressRepository instance.
func New(conn database.DB) domain.AddressRepository {
	return &AddressRepository{
		conn,
	}
}

// Create stores a new Address.
func (ar *AddressRepository) Create(ctx context.Context, addr *domain.Address) error {
	dao := ar.toDAO(addr)

	result := ar.conn.DB().Create(&dao).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "duplicated key not allowed" {
			return helper.ErrDataAlreadyExists
		}

		return result.Error
	}

	return nil
}

// GetAll returns all Addresses.
func (ar *AddressRepository) GetAll(ctx context.Context, userID uint, title string) ([]domain.Address, error) {
	var (
		addrs []domain.Address
		daos  []dao.Address
		db    = ar.conn.DB()
	)

	if title != "" {
		db = ar.conn.DB().Where("title LIKE ?", "%"+title+"%")
	}

	result := db.Find(&daos).Where("user_id = ?", userID).WithContext(ctx)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, dao := range daos {
		addrs = append(addrs, *ar.toDomain(&dao))
	}

	return addrs, nil
}

// GetByID returns the Address with the specified ID.
func (ar *AddressRepository) GetByID(ctx context.Context, userID, id uint) (*domain.Address, error) {
	var dao dao.Address

	result := ar.conn.DB().First(&dao, id).Where("user_id = ?", userID).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, helper.ErrDataNotFound
		}

		return nil, result.Error
	}

	addr := ar.toDomain(&dao)

	return addr, nil
}

// Update updates the Address with the specified ID.
func (ar *AddressRepository) Update(ctx context.Context, addr *domain.Address) error {
	dao := ar.toDAO(addr)

	result := ar.conn.DB().Model(dao).Updates(&dao).Where("id = ?", dao.ID).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return helper.ErrDataNotFound
		}

		return result.Error
	}

	return nil
}

// Delete removes the Address with the specified ID.
func (ar *AddressRepository) Delete(ctx context.Context, userID, id uint) error {
	var dao dao.Address

	result := ar.conn.DB().Delete(&dao, id).Where("user_id = ?", userID).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return helper.ErrDataNotFound
		}

		return result.Error
	}

	return nil
}

// toDomain converts a DAO Address to a Address.
func (ar *AddressRepository) toDomain(addr *dao.Address) *domain.Address {
	return &domain.Address{
		ID:          addr.ID,
		Title:       addr.Title,
		Receiver:    addr.Receiver,
		PhoneNumber: addr.PhoneNumber,
		Details:     addr.Details,
		UserID:      addr.UserID,
	}
}

// toDAO converts a Address to a DAO Address.
func (ur *AddressRepository) toDAO(addr *domain.Address) *dao.Address {
	return &dao.Address{
		Model: dao.Model{
			ID: addr.ID,
		},
		Title:       addr.Title,
		Receiver:    addr.Receiver,
		PhoneNumber: addr.PhoneNumber,
		Details:     addr.Details,
		UserID:      addr.UserID,
	}
}
