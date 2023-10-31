package auth

import (
	v1 "github.com/bagashiz/gommerce/internal/app/auth/controller/http/v1"
	"github.com/bagashiz/gommerce/internal/app/auth/repository"
	"github.com/bagashiz/gommerce/internal/app/auth/usecase"
	cityRepository "github.com/bagashiz/gommerce/internal/app/city/repository"
	provinceRepository "github.com/bagashiz/gommerce/internal/app/province/repository"
	userRepository "github.com/bagashiz/gommerce/internal/app/user/repository"
	"github.com/bagashiz/gommerce/internal/pkg/database"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
	"github.com/bagashiz/gommerce/internal/pkg/token"
)

// New injects the dependencies of auth package
func New(db database.DB, server *http.Http, token token.Token) {
	authRepo := repository.New(db)
	userRepo := userRepository.New(db)
	cityRepo := cityRepository.New()
	provRepo := provinceRepository.New()
	authUsecase := usecase.New(authRepo, userRepo, cityRepo, provRepo, token)
	authV1 := v1.New(authUsecase, server)

	routeV1 := server.App.Group("/v1/auth")
	authV1.InitRoutes(routeV1)
}
