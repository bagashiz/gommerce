package province

import (
	v1 "github.com/bagashiz/gommerce/internal/app/province/controller/http/v1"
	"github.com/bagashiz/gommerce/internal/app/province/repository"
	"github.com/bagashiz/gommerce/internal/app/province/usecase"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
)

// New injects the dependencies of province package
func New(server *http.Http) {
	provinceRepo := repository.New()
	provinceUsecase := usecase.New(provinceRepo)
	provinceV1 := v1.New(provinceUsecase, server)

	routeV1 := server.App.Group("/v1/provinces")
	provinceV1.InitRoutes(routeV1)
}
