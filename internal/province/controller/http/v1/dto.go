package v1

import "github.com/bagashiz/gommerce/internal/province/domain"

// provinceParam is a struct for validating province path parameter
type provinceParam struct {
	ID string `params:"id" validate:"required"`
}

// provinceResponse is a struct for structuring province response
type provinceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewProvinceResponse creates a new instance of ProvinceResponse
func NewProvinceResponse(province domain.Province) *provinceResponse {
	return &provinceResponse{
		ID:   province.ID,
		Name: province.Name,
	}
}
