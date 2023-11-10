package v1

import (
	cityDom "github.com/bagashiz/gommerce/internal/app/city/domain"
	provDom "github.com/bagashiz/gommerce/internal/app/province/domain"
	"github.com/bagashiz/gommerce/internal/app/user/domain"
)

// updateUserRequest is a struct for validating update user request body
type updateUserRequest struct {
	Name        string `json:"name" validate:"omitempty,required"`
	Password    string `json:"password" validate:"omitempty,required"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,required"`
	Email       string `json:"email" validate:"omitempty,required"`
	BirthDate   string `json:"birth_date" validate:"omitempty,required"`
	About       string `json:"about" validate:"omitempty,required"`
	Job         string `json:"job" validate:"omitempty,required"`
	ProvinceID  string `json:"province_id" validate:"omitempty,required"`
	CityID      string `json:"city_id" validate:"omitempty,required"`
}

// userResponse is a struct for structuring user response
type userResponse struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	PhoneNumber string            `json:"phone_number"`
	Email       string            `json:"email"`
	BirthDate   string            `json:"birth_date" `
	About       string            `json:"about"`
	Job         string            `json:"job"`
	Province    *provDom.Province `json:"province"`
	City        *cityDom.City     `json:"city"`
}

// NewUserResponse creates a new instance of UserResponse
func NewUserResponse(user domain.User) *userResponse {
	return &userResponse{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		BirthDate:   user.BirthDate.Format("02-01-2006"),
		About:       user.About,
		Job:         user.Job,
		Province:    user.Province,
		City:        user.City,
	}
}
