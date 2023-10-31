package repository

import (
	"context"
	"strings"

	"github.com/bagashiz/gommerce/internal/app/auth/domain"
	userDom "github.com/bagashiz/gommerce/internal/app/user/domain"
	"github.com/bagashiz/gommerce/internal/pkg/database"
	"github.com/bagashiz/gommerce/internal/pkg/database/dao"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
)

// AuthRepository is a struct that implements AuthRepository interface.
type AuthRepository struct {
	conn database.DB
}

// New creates a new AuthRepository instance.
func New(conn database.DB) domain.AuthRepository {
	return &AuthRepository{
		conn,
	}
}

// Create stores a new User and related Shop.
func (ar *AuthRepository) Create(ctx context.Context, user *userDom.User) error {
	tx := ar.conn.DB().Begin().WithContext(ctx)

	// create user
	userDAO := ar.toDAO(user)

	resultUser := tx.Model(userDAO).Create(&userDAO).WithContext(ctx)
	if resultUser.Error != nil {
		tx.Rollback()

		if resultUser.Error.Error() == "duplicated key not allowed" {
			return helper.ErrDataAlreadyExists
		}

		return resultUser.Error
	}

	// create shop
	shopName := strings.ToLower(strings.ReplaceAll(user.Name, " ", "-")) + "-shop"

	shopDAO := &dao.Shop{
		Name:   shopName,
		UserID: userDAO.ID,
	}

	resultShop := tx.Model(shopDAO).Create(&shopDAO).WithContext(ctx)
	if resultShop.Error != nil {
		tx.Rollback()

		if resultShop.Error.Error() == "duplicated key not allowed" {
			return helper.ErrDataAlreadyExists
		}

		return resultShop.Error
	}

	tx.Commit()

	return nil
}

// GetByPhoneNumber returns the User with the specified phone number.
func (ar *AuthRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*userDom.User, error) {
	var dao dao.User

	result := ar.conn.DB().First(&dao).Where("phone_number = ?", phoneNumber).WithContext(ctx)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, helper.ErrDataNotFound
		}

		return nil, result.Error
	}

	user := ar.toDomain(&dao)

	return user, nil
}

// toDomain converts a DAO Category to a Category.
func (ar *AuthRepository) toDomain(user *dao.User) *userDom.User {
	return &userDom.User{
		ID:          user.ID,
		Name:        user.Name,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		BirthDate:   user.BirthDate,
		About:       user.About,
		Job:         user.Job,
		ProvinceID:  user.ProvinceID,
		CityID:      user.CityID,
	}
}

// toDAO converts a Category to a DAO Category.
func (ar *AuthRepository) toDAO(user *userDom.User) *dao.User {
	return &dao.User{
		Model: dao.Model{
			ID: user.ID,
		},
		Name:        user.Name,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		BirthDate:   user.BirthDate,
		About:       user.About,
		Job:         user.Job,
		ProvinceID:  user.ProvinceID,
		CityID:      user.CityID,
	}
}
