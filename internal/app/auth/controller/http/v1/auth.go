package v1

import (
	"errors"

	"github.com/bagashiz/gommerce/internal/app/auth/domain"
	userDom "github.com/bagashiz/gommerce/internal/app/user/domain"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
	"github.com/gofiber/fiber/v2"
)

// AuthControllerV1 is a struct for version 1 of AuthController
type AuthControllerV1 struct {
	uc     domain.AuthUsecase
	server *http.Http
}

// New creates a new instance of AuthControllerV1
func New(uc domain.AuthUsecase, server *http.Http) *AuthControllerV1 {
	return &AuthControllerV1{
		uc:     uc,
		server: server,
	}
}

// Register handles POST /auth/register request
func (ac *AuthControllerV1) Register(ctx *fiber.Ctx) error {
	var req registerRequest

	if err := ctx.BodyParser(&req); err != nil {
		ac.server.Logger.Error("failed to parse request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPOSTDATA, err, nil)
	}

	birthDate, err := helper.ParseTime(req.BirthDate)
	if err != nil {
		ac.server.Logger.Error("failed to parse request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPOSTDATA, err, nil)
	}

	user := userDom.User{
		Name:        req.Name,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		BirthDate:   birthDate,
		Job:         req.Job,
		ProvinceID:  req.ProvinceID,
		CityID:      req.CityID,
	}

	if err := ac.uc.Register(ctx.Context(), &user); err != nil {
		ac.server.Logger.Error("failed to register user", "error", err)

		if errors.Is(err, helper.ErrDataAlreadyExists) {
			return helper.Response(ctx, fiber.StatusConflict, false, helper.FAILEDPOSTDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDPOSTDATA, err, nil)
	}

	return helper.Response(ctx, fiber.StatusCreated, true, helper.SUCCEEDPOSTDATA, nil, nil)
}

// Login handles POST /auth/login request
func (ac *AuthControllerV1) Login(ctx *fiber.Ctx) error {
	var req loginRequest

	if err := ctx.BodyParser(&req); err != nil {
		ac.server.Logger.Error("failed to parse request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPOSTDATA, err, nil)
	}

	user, city, province, token, err := ac.uc.Login(ctx.Context(), req.PhoneNumber, req.Password)
	if err != nil {
		ac.server.Logger.Error("failed to login user", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDPOSTDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDPOSTDATA, err, nil)
	}

	rsp := newLoginResponse(user, city, province, token)

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDPOSTDATA, nil, rsp)
}
