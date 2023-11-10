package user

import (
	cityRepository "github.com/bagashiz/gommerce/internal/app/city/repository"
	provinceRepository "github.com/bagashiz/gommerce/internal/app/province/repository"
	v1 "github.com/bagashiz/gommerce/internal/app/user/controller/http/v1"
	"github.com/bagashiz/gommerce/internal/app/user/repository"
	"github.com/bagashiz/gommerce/internal/app/user/usecase"
	"github.com/bagashiz/gommerce/internal/pkg/database"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
	"github.com/bagashiz/gommerce/internal/pkg/token"
)

// New injects the dependencies of user package
func New(db database.DB, server *http.Http, token token.Token) {
	userRepo := repository.New(db)
	cityRepo := cityRepository.New()
	provRepo := provinceRepository.New()
	userUsecase := usecase.New(userRepo, cityRepo, provRepo)
	userV1 := v1.New(userUsecase, server)

	routeV1 := server.App.Group("/v1/users")
	userV1.InitRoutes(routeV1, token)
}
