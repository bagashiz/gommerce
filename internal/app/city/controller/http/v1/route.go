package v1

import "github.com/gofiber/fiber/v2"

// InitRoutes registers all routes for version 1
func (cc *CityControllerV1) InitRoutes(routeV1 fiber.Router) {
	routeV1.Get("/:prov_id/cities/:city_id", cc.GetByID)
	routeV1.Get("/:prov_id/cities", cc.GetAll)
}
