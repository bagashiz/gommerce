package v1

import (
	"github.com/bagashiz/gommerce/internal/pkg/middleware"
	"github.com/bagashiz/gommerce/internal/pkg/token"
	"github.com/gofiber/fiber/v2"
)

// InitRoutes registers all routes for version 1
func (cc *CategoryControllerV1) InitRoutes(routeV1 fiber.Router, token token.Token) {
	routeV1.Get("/", cc.GetAll)
	routeV1.Post("/", middleware.AuthMiddleware(token), middleware.AdminMiddleware(), cc.Create)
	routeV1.Get("/:id", middleware.AuthMiddleware(token), cc.GetByID)
	routeV1.Put("/:id", middleware.AuthMiddleware(token), middleware.AdminMiddleware(), cc.Update)
	routeV1.Delete("/:id", middleware.AuthMiddleware(token), middleware.AdminMiddleware(), cc.Delete)
}
