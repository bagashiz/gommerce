package category

import (
	v1 "github.com/bagashiz/gommerce/internal/app/category/controller/http/v1"
	"github.com/bagashiz/gommerce/internal/app/category/repository"
	"github.com/bagashiz/gommerce/internal/app/category/usecase"
	"github.com/bagashiz/gommerce/internal/pkg/database"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
	"github.com/bagashiz/gommerce/internal/pkg/token"
)

// New injects the dependencies of category package
func New(db database.DB, server *http.Http, token token.Token) {
	categoryRepo := repository.New(db)
	categoryUsecase := usecase.New(categoryRepo)
	categoryV1 := v1.New(categoryUsecase, server)

	routeV1 := server.App.Group("/v1/categories")
	categoryV1.InitRoutes(routeV1, token)
}
