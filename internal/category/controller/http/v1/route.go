package v1

import "github.com/gofiber/fiber/v2"

// InitRoutes registers all routes for version 1
func (cc *CategoryControllerV1) InitRoutes(routeV1 fiber.Router) {

	routeV1.Post("/", cc.Create)
	routeV1.Get("/", cc.GetAll)
	routeV1.Get("/:id", cc.GetByID)
	routeV1.Put("/:id", cc.Update)
	routeV1.Delete("/:id", cc.Delete)
}
