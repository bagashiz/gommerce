package v1

import (
	"errors"

	"github.com/bagashiz/gommerce/internal/app/shop/domain"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
	"github.com/gofiber/fiber/v2"
)

// ShopControllerV1 is a struct for version 1 of ShopController
type ShopControllerV1 struct {
	uc     domain.ShopUsecase
	server *http.Http
}

// New creates a new instance of ShopControllerV1
func New(uc domain.ShopUsecase, server *http.Http) *ShopControllerV1 {
	return &ShopControllerV1{
		uc,
		server,
	}
}

// GetAll handles GET /shops request
func (sc *ShopControllerV1) GetAll(ctx *fiber.Ctx) error {
	var q shopQuery

	if err := ctx.QueryParser(&q); err != nil {
		sc.server.Logger.Error("failed to parse query", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	if err := sc.server.Validate.Struct(&q); err != nil {
		sc.server.Logger.Error("failed to validate query", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	res, err := sc.uc.GetAll(ctx.Context(), q.Page, q.Limit, q.Name)
	if err != nil {
		sc.server.Logger.Error("failed to get shops", "error", err)
		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	var shops []shopResponse

	for _, shop := range res {
		shops = append(shops, *NewShopResponse(shop))
	}

	rsp := &shopResponsePaginated{
		Shops: shops,
		Page:  q.Page,
		Limit: q.Limit,
	}

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}

// GetUserShop handles GET /shops/my request
func (sc *ShopControllerV1) GetUserShop(ctx *fiber.Ctx) error {
	idCtx := ctx.Locals("id") // TODO: implement middleware to get user id from token
	if idCtx == nil {
		sc.server.Logger.Error("failed to get user id from context")
		return helper.Response(ctx, fiber.StatusUnauthorized, false, helper.FAILEDGETDATA, nil, nil)
	}

	userID := idCtx.(uint)

	res, err := sc.uc.GetUserShop(ctx.Context(), userID)
	if err != nil {
		sc.server.Logger.Error("failed to get shop", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	rsp := NewShopResponse(*res)

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}

// GetByID handles GET /shops/:id request
func (sc *ShopControllerV1) GetByID(ctx *fiber.Ctx) error {
	var p shopParam

	if err := ctx.ParamsParser(&p); err != nil {
		sc.server.Logger.Error("failed to parse path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	if err := sc.server.Validate.Struct(&p); err != nil {
		sc.server.Logger.Error("failed to validate path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	res, err := sc.uc.GetByID(ctx.Context(), p.ID)
	if err != nil {
		sc.server.Logger.Error("failed to get shop", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	rsp := NewShopResponse(*res)

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}

// Update handles PUT /shops/:id request
func (sc *ShopControllerV1) Update(ctx *fiber.Ctx) error {
	idCtx := ctx.Locals("id") // TODO: implement middleware to get user id from token
	if idCtx == nil {
		sc.server.Logger.Error("failed to get user id from context")
		return helper.Response(ctx, fiber.StatusUnauthorized, false, helper.FAILEDPUTDATA, nil, nil)
	}

	userID := idCtx.(uint)

	var req updateShopRequest

	if err := ctx.BodyParser(&req); err != nil {
		sc.server.Logger.Error("failed to parse request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	if err := sc.server.Validate.Struct(&req); err != nil {
		sc.server.Logger.Error("failed to validate request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	var p shopParam

	if err := ctx.ParamsParser(&p); err != nil {
		sc.server.Logger.Error("failed to parse path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	if err := sc.server.Validate.Struct(&p); err != nil {
		sc.server.Logger.Error("failed to validate path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	if err := sc.uc.Update(ctx.Context(), userID, p.ID); err != nil {
		sc.server.Logger.Error("failed to update shop", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}
		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDPUTDATA, err, nil)
	}

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDPUTDATA, nil, nil)
}
