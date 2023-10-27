package v1

import (
	"github.com/bagashiz/gommerce/internal/app/city/domain"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
	"github.com/gofiber/fiber/v2"
)

// CityControllerV1 is a struct for version 1 of CityController
type CityControllerV1 struct {
	uc     domain.CityUsecase
	server *http.Http
}

// New creates a new instance of CityControllerV1
func New(uc domain.CityUsecase, server *http.Http) *CityControllerV1 {
	return &CityControllerV1{
		uc,
		server,
	}
}

// GetAll is a function to get all cities
func (cc *CityControllerV1) GetAll(ctx *fiber.Ctx) error {
	var p citiesParam

	if err := ctx.ParamsParser(&p); err != nil {
		cc.server.Logger.Error("failed to parse path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	res, err := cc.uc.GetAll(p.ProvinceID)
	if err != nil {
		cc.server.Logger.Error("failed to get cities", "error", err)

		if err == helper.ErrDataNotFound {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	var rsp []cityResponse

	for _, city := range res {
		rsp = append(rsp, *NewCityResponse(city))
	}

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}

// GetByID is a function to get city by id
func (cc *CityControllerV1) GetByID(ctx *fiber.Ctx) error {
	var p cityParam

	if err := ctx.ParamsParser(&p); err != nil {
		cc.server.Logger.Error("failed to parse path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	res, err := cc.uc.GetByID(p.ProvinceID, p.ID)
	if err != nil {
		cc.server.Logger.Error("failed to get city", "error", err)

		if err == helper.ErrDataNotFound {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	rsp := NewCityResponse(*res)

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}
