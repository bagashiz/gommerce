package v1

import "github.com/gofiber/fiber/v2"

// InitRoutes registers all routes for version 1
func (ac *AddressControllerV1) InitRoutes(routeV1 fiber.Router) {
	routeV1.Post("/", ac.Create)
	routeV1.Get("/", ac.GetAll)
	routeV1.Get("/:id", ac.GetByID)
	routeV1.Put("/:id", ac.Update)
	routeV1.Delete("/:id", ac.Delete)
}
