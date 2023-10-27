package v1

import (
	"github.com/bagashiz/gommerce/internal/pkg/helper"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
	"github.com/bagashiz/gommerce/internal/province/domain"
	"github.com/gofiber/fiber/v2"
)

// ProvinceControllerV1 is a struct for version 1 of ProvinceController
type ProvinceControllerV1 struct {
	uc     domain.ProvinceUsecase
	server *http.Http
}

// New creates a new instance of ProvinceControllerV1
func New(uc domain.ProvinceUsecase, server *http.Http) *ProvinceControllerV1 {
	return &ProvinceControllerV1{
		uc,
		server,
	}
}

// GetAll is a function to get all provinces
func (pc *ProvinceControllerV1) GetAll(ctx *fiber.Ctx) error {
	res, err := pc.uc.GetAll()
	if err != nil {
		pc.server.Logger.Error("failed to get provinces", "error", err)

		if err == helper.ErrDataNotFound {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	var rsp []provinceResponse

	for _, province := range res {
		rsp = append(rsp, *NewProvinceResponse(province))
	}

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}

// GetByID is a function to get province by id
func (pc *ProvinceControllerV1) GetByID(ctx *fiber.Ctx) error {
	var p provinceParam

	if err := ctx.ParamsParser(&p); err != nil {
		pc.server.Logger.Error("failed to parse path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	res, err := pc.uc.GetByID(p.ID)
	if err != nil {
		pc.server.Logger.Error("failed to get province", "error", err)

		if err == helper.ErrDataNotFound {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	rsp := NewProvinceResponse(*res)

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}
