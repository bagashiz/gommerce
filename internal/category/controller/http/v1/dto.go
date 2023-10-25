package v1

import "github.com/bagashiz/gommerce/internal/category/domain"

// categoryParam is a struct for validating category path parameter
type categoryParam struct {
	ID uint `uri:"id" validate:"required"`
}

// categoryQuery is a struct for validating category query parameter
type categoryQuery struct {
	Page  int `form:"page" validate:"required,min=1"`
	Limit int `form:"limit" validate:"required,min=5"`
}

// createCategoryRequest is a struct for validating create category request body
type createCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

// updateCategoryRequest is a struct for validating update category request body
type updateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

// categoryResponse is a struct for structuring category response
type categoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// categoryResponsePaginated is a struct for structuring paginated category response
type categoryResponsePaginated struct {
	Categories []categoryResponse `json:"categories"`
	Page       int                `json:"page"`
	Limit      int                `json:"limit"`
}

// NewCategoryResponse creates a new instance of CategoryResponse
func NewCategoryResponse(category domain.Category) *categoryResponse {
	return &categoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}
