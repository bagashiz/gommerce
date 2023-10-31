package v1

import (
	cityDom "github.com/bagashiz/gommerce/internal/app/city/domain"
	provinceDom "github.com/bagashiz/gommerce/internal/app/province/domain"
	userDom "github.com/bagashiz/gommerce/internal/app/user/domain"
)

// loginRequest is a struct for validating user login request body.
type loginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required,min=8"`
}

// registerRequest is a struct for validating user register request body.
type registerRequest struct {
	Name        string `json:"name" validate:"required"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"required"`
	BirthDate   string `json:"birth_date" validate:"required"`
	Job         string `json:"job" validate:"required"`
	ProvinceID  string `json:"province_id" validate:"required"`
	CityID      string `json:"city_id" validate:"required"`
}

// loginResponse is a struct for returning user login response body.
type loginResponse struct {
	Name        string           `json:"name"`
	PhoneNumber string           `json:"phone_number"`
	Email       string           `json:"email"`
	BirthDate   string           `json:"birth_date"`
	Job         string           `json:"job"`
	Province    provinceResponse `json:"province"`
	City        cityResponse     `json:"city"`
	Token       string           `json:"token"`
}

// provinceResponse is a struct for structuring province data.
type provinceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// cityResponse is a struct for structuring city data.
type cityResponse struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

// newLoginResponse is a function to create loginResponse struct.
func newLoginResponse(user *userDom.User, city *cityDom.City, province *provinceDom.Province, token string) *loginResponse {
	return &loginResponse{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		BirthDate:   user.BirthDate.Format("02-01-2006"),
		Job:         user.Job,
		Province: provinceResponse{
			ID:   province.ID,
			Name: province.Name,
		},
		City: cityResponse{
			ID:         city.ID,
			ProvinceID: city.ProvinceID,
			Name:       city.Name,
		},
		Token: token,
	}
}
