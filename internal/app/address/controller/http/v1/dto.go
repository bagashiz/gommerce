package v1

import "github.com/bagashiz/gommerce/internal/app/address/domain"

// addrParam is a struct for validating address path parameter
type addrParam struct {
	ID uint `params:"id" validate:"required"`
}

// addrQuery is a struct for validating address query parameter
type addrQuery struct {
	Title string `form:"title" validate:"omitempty"`
}

// createAddrRequest is a struct for validating create address request body
type createAddrRequest struct {
	Title       string `json:"title" validate:"required"`
	Receiver    string `json:"receiver" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Details     string `json:"details" validate:"required"`
}

// updateAddrRequest is a struct for validating update address request body
type updateAddrRequest struct {
	Title       string `json:"title" validate:"omitempty,required"`
	Receiver    string `json:"receiver" validate:"omitempty,required"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,required"`
	Details     string `json:"details" validate:"omitempty,required"`
}

// addrResponse is a struct for structuring address response
type addrResponse struct {
	Title       string `json:"title"`
	Receiver    string `json:"receiver"`
	PhoneNumber string `json:"phone_number"`
	Details     string `json:"details"`
}

// NewAddrResponse creates a new instance of AddrResponse
func NewAddrResponse(addr domain.Address) *addrResponse {
	return &addrResponse{
		Title:       addr.Title,
		Receiver:    addr.Receiver,
		PhoneNumber: addr.PhoneNumber,
		Details:     addr.Details,
	}
}
