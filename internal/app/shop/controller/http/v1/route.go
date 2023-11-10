package v1

import (
	"github.com/bagashiz/gommerce/internal/pkg/middleware"
	"github.com/bagashiz/gommerce/internal/pkg/token"
	"github.com/gofiber/fiber/v2"
)

// InitRoutes registers all routes for version 1
func (sc *ShopControllerV1) InitRoutes(routeV1 fiber.Router, token token.Token) {
	routeV1.Get("/my", middleware.AuthMiddleware(token), sc.GetUserShop)
	routeV1.Get("/", sc.GetAll)
	routeV1.Get("/:id", sc.GetByID)
	routeV1.Put("/:id", middleware.AuthMiddleware(token), sc.Update)
}
