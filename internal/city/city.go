package city

import (
	v1 "github.com/bagashiz/gommerce/internal/city/controller/http/v1"
	"github.com/bagashiz/gommerce/internal/city/repository"
	"github.com/bagashiz/gommerce/internal/city/usecase"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
)

// New injects the dependencies of province package
func New(server *http.Http) {
	cityRepo := repository.New()
	cityUsecase := usecase.New(cityRepo)
	cityV1 := v1.New(cityUsecase, server)

	routeV1 := server.App.Group("/v1/provinces")
	cityV1.InitRoutes(routeV1)
}
