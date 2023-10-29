package shop

import (
	v1 "github.com/bagashiz/gommerce/internal/app/shop/controller/http/v1"
	"github.com/bagashiz/gommerce/internal/app/shop/repository"
	"github.com/bagashiz/gommerce/internal/app/shop/usecase"
	"github.com/bagashiz/gommerce/internal/pkg/database"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
)

// New injects the dependencies of shop package
func New(db database.DB, server *http.Http) {
	shopRepo := repository.New(db)
	shopUsecase := usecase.New(shopRepo)
	shopV1 := v1.New(shopUsecase, server)

	routeV1 := server.App.Group("/v1/shops")
	shopV1.InitRoutes(routeV1)
}
