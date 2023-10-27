package v1

import "github.com/gofiber/fiber/v2"

// InitRoutes registers all routes for version 1
func (pc *ProvinceControllerV1) InitRoutes(routeV1 fiber.Router) {
	routeV1.Get("/:id", pc.GetByID)
	routeV1.Get("/", pc.GetAll)
}
