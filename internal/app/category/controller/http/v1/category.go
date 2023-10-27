package v1

import (
	"errors"

	"github.com/bagashiz/gommerce/internal/app/category/domain"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
	"github.com/gofiber/fiber/v2"
)

// CategoryControllerV1 is a struct for version 1 of CategoryController
type CategoryControllerV1 struct {
	uc     domain.CategoryUsecase
	server *http.Http
}

// New creates a new instance of CategoryControllerV1
func New(uc domain.CategoryUsecase, server *http.Http) *CategoryControllerV1 {
	return &CategoryControllerV1{
		uc,
		server,
	}
}

// Create handles POST /categories request
func (cc *CategoryControllerV1) Create(ctx *fiber.Ctx) error {
	var req createCategoryRequest

	if err := ctx.BodyParser(&req); err != nil {
		cc.server.Logger.Error("failed to parse request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPOSTDATA, err, nil)
	}

	if err := cc.server.Validate.Struct(&req); err != nil {
		cc.server.Logger.Error("failed to validate request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPOSTDATA, err, nil)
	}

	category := &domain.Category{
		Name: req.Name,
	}

	if err := cc.uc.Create(ctx.Context(), category); err != nil {
		cc.server.Logger.Error("failed to create category", "error", err)

		if errors.Is(err, helper.ErrDataAlreadyExists) {
			return helper.Response(ctx, fiber.StatusConflict, false, helper.FAILEDPOSTDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDPOSTDATA, err, nil)
	}

	return helper.Response(ctx, fiber.StatusCreated, true, helper.SUCCEEDPOSTDATA, nil, nil)
}

// GetAll handles GET /categories request
func (cc *CategoryControllerV1) GetAll(ctx *fiber.Ctx) error {
	var q categoryQuery

	if err := ctx.QueryParser(&q); err != nil {
		cc.server.Logger.Error("failed to parse query", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	if err := cc.server.Validate.Struct(&q); err != nil {
		cc.server.Logger.Error("failed to validate query", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	res, err := cc.uc.GetAll(ctx.Context(), q.Page, q.Limit)
	if err != nil {
		cc.server.Logger.Error("failed to get categories", "error", err)
		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	var categories []categoryResponse

	for _, category := range res {
		categories = append(categories, *NewCategoryResponse(category))
	}

	rsp := &categoryResponsePaginated{
		Categories: categories,
		Page:       q.Page,
		Limit:      q.Limit,
	}

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}

// GetByID handles GET /categories/:id request
func (cc *CategoryControllerV1) GetByID(ctx *fiber.Ctx) error {
	var p categoryParam

	if err := ctx.ParamsParser(&p); err != nil {
		cc.server.Logger.Error("failed to parse path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	if err := cc.server.Validate.Struct(&p); err != nil {
		cc.server.Logger.Error("failed to validate path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	res, err := cc.uc.GetByID(ctx.Context(), p.ID)
	if err != nil {
		cc.server.Logger.Error("failed to get category", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	rsp := NewCategoryResponse(*res)

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}

// Update handles PUT /categories/:id request
func (cc *CategoryControllerV1) Update(ctx *fiber.Ctx) error {
	var req updateCategoryRequest

	if err := ctx.BodyParser(&req); err != nil {
		cc.server.Logger.Error("failed to parse request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	if err := cc.server.Validate.Struct(&req); err != nil {
		cc.server.Logger.Error("failed to validate request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	var p categoryParam

	if err := ctx.ParamsParser(&p); err != nil {
		cc.server.Logger.Error("failed to parse path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	if err := cc.server.Validate.Struct(&p); err != nil {
		cc.server.Logger.Error("failed to validate path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	category := &domain.Category{
		ID:   p.ID,
		Name: req.Name,
	}

	if err := cc.uc.Update(ctx.Context(), category); err != nil {
		cc.server.Logger.Error("failed to update category", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}
		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDPUTDATA, err, nil)
	}

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDPUTDATA, nil, nil)
}

// Delete handles DELETE /categories/:id request
func (cc *CategoryControllerV1) Delete(ctx *fiber.Ctx) error {
	var p categoryParam

	if err := ctx.ParamsParser(&p); err != nil {
		cc.server.Logger.Error("failed to parse path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDDELETEDATA, err, nil)
	}

	if err := cc.server.Validate.Struct(&p); err != nil {
		cc.server.Logger.Error("failed to validate path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDDELETEDATA, err, nil)
	}

	if err := cc.uc.Delete(ctx.Context(), p.ID); err != nil {
		cc.server.Logger.Error("failed to delete category", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDDELETEDATA, err, nil)
	}

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDDELETEDATA, nil, nil)
}
