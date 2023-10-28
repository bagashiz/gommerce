package v1

import "github.com/gofiber/fiber/v2"

// InitRoutes registers all routes for version 1
func (uc *UserControllerV1) InitRoutes(routeV1 fiber.Router) {
	routeV1.Get("/:id", uc.GetByID)
	routeV1.Put("/:id", uc.Update)
}
