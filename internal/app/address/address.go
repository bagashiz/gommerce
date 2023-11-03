package address

import (
	v1 "github.com/bagashiz/gommerce/internal/app/address/controller/http/v1"
	"github.com/bagashiz/gommerce/internal/app/address/repository"
	"github.com/bagashiz/gommerce/internal/app/address/usecase"
	"github.com/bagashiz/gommerce/internal/pkg/database"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
)

// New injects the dependencies of address package
func New(db database.DB, server *http.Http) {
	addrRepo := repository.New(db)
	addrUsecase := usecase.New(addrRepo)
	addrV1 := v1.New(addrUsecase, server)

	routeV1 := server.App.Group("/v1/users/addresses")
	addrV1.InitRoutes(routeV1)
}
