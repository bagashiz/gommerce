package helper

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Response is a struct for structuring HTTP response in JSON format.
type JSONResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
	Data    any      `json:"data"`
}

// Response creates a new instance of JSONResponse.
func Response(ctx *fiber.Ctx, code int, status bool, message string, err error, data any) error {
	var errMsgs []string

	if err != nil {
		errMsgs = parseError(err)
	}

	res := &JSONResponse{
		Status:  status,
		Message: message,
		Errors:  errMsgs,
		Data:    data,
	}

	return ctx.Status(code).JSON(res)
}

// parseError parses error messages from error object to a slice of string.
func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}
