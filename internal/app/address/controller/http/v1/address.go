package v1

import (
	"errors"

	"github.com/bagashiz/gommerce/internal/app/address/domain"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
	"github.com/gofiber/fiber/v2"
)

// AddressControllerV1 is a struct for version 1 of AddressController
type AddressControllerV1 struct {
	uc     domain.AddressUsecase
	server *http.Http
}

// NewAddressControllerV1 creates a new instance of AddressControllerV1
func New(uc domain.AddressUsecase, server *http.Http) *AddressControllerV1 {
	return &AddressControllerV1{
		uc:     uc,
		server: server,
	}
}

// Create handles POST /users/Addresses request
func (ac *AddressControllerV1) Create(ctx *fiber.Ctx) error {
	var req createAddrRequest

	if err := ctx.BodyParser(&req); err != nil {
		ac.server.Logger.Error("failed to parse request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPOSTDATA, err, nil)
	}

	if err := ac.server.Validate.Struct(&req); err != nil {
		ac.server.Logger.Error("failed to validate request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPOSTDATA, err, nil)
	}

	userID := ctx.Locals("user_id").(uint)

	addr := &domain.Address{
		Title:       req.Title,
		Receiver:    req.Receiver,
		PhoneNumber: req.PhoneNumber,
		Details:     req.Details,
		UserID:      userID,
	}

	if err := ac.uc.Create(ctx.Context(), addr); err != nil {
		ac.server.Logger.Error("failed to create address", "error", err)
		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDPOSTDATA, err, nil)
	}

	return helper.Response(ctx, fiber.StatusCreated, true, helper.SUCCEEDPOSTDATA, nil, nil)
}

// GetAll handles GET /users/Addresses request
func (ac *AddressControllerV1) GetAll(ctx *fiber.Ctx) error {
	var q addrQuery

	if err := ctx.QueryParser(&q); err != nil {
		ac.server.Logger.Error("failed to parse query parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	if err := ac.server.Validate.Struct(&q); err != nil {
		ac.server.Logger.Error("failed to validate query parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	userID := ctx.Locals("user_id").(uint)

	res, err := ac.uc.GetAll(ctx.Context(), userID, q.Title)
	if err != nil {
		ac.server.Logger.Error("failed to get addresses", "error", err)
		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	rsp := make([]*addrResponse, 0, len(res))

	for _, addr := range res {
		rsp = append(rsp, NewAddrResponse(addr))
	}

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}

// GetByID handles GET /users/Addresses request
func (ac *AddressControllerV1) GetByID(ctx *fiber.Ctx) error {
	var p addrParam

	if err := ctx.ParamsParser(&p); err != nil {
		ac.server.Logger.Error("failed to parse path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	if err := ac.server.Validate.Struct(&p); err != nil {
		ac.server.Logger.Error("failed to validate path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	userID := ctx.Locals("user_id").(uint)

	res, err := ac.uc.GetByID(ctx.Context(), userID, p.ID)
	if err != nil {
		ac.server.Logger.Error("failed to get address", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	rsp := NewAddrResponse(*res)

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}

// Update handles PUT /users/Addresses request
func (ac *AddressControllerV1) Update(ctx *fiber.Ctx) error {
	var req updateAddrRequest

	if err := ctx.BodyParser(&req); err != nil {
		ac.server.Logger.Error("failed to parse request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	if err := ac.server.Validate.Struct(&req); err != nil {
		ac.server.Logger.Error("failed to validate request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	var p addrParam

	if err := ctx.ParamsParser(&p); err != nil {
		ac.server.Logger.Error("failed to parse path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	if err := ac.server.Validate.Struct(&p); err != nil {
		ac.server.Logger.Error("failed to validate path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	userID := ctx.Locals("user_id").(uint)

	addr := &domain.Address{
		ID:          p.ID,
		Title:       req.Title,
		Receiver:    req.Receiver,
		PhoneNumber: req.PhoneNumber,
		Details:     req.Details,
		UserID:      userID,
	}

	if err := ac.uc.Update(ctx.Context(), addr); err != nil {
		ac.server.Logger.Error("failed to update address", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDPUTDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDPUTDATA, err, nil)
	}

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDPUTDATA, nil, nil)
}

// Delete handles DELETE /users/Addresses request
func (ac *AddressControllerV1) Delete(ctx *fiber.Ctx) error {
	var p addrParam

	if err := ctx.ParamsParser(&p); err != nil {
		ac.server.Logger.Error("failed to parse path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDDELETEDATA, err, nil)
	}

	if err := ac.server.Validate.Struct(&p); err != nil {
		ac.server.Logger.Error("failed to validate path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDDELETEDATA, err, nil)
	}

	userID := ctx.Locals("user_id").(uint)

	if err := ac.uc.Delete(ctx.Context(), userID, p.ID); err != nil {
		ac.server.Logger.Error("failed to delete address", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDDELETEDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDDELETEDATA, err, nil)
	}

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDDELETEDATA, nil, nil)
}
