package v1

import "github.com/gofiber/fiber/v2"

// InitRoutes registers all routes for version 1
func (sc *ShopControllerV1) InitRoutes(routeV1 fiber.Router) {
	routeV1.Get("/my", sc.GetUserShop)
	routeV1.Get("/", sc.GetAll)
	routeV1.Get("/:id", sc.GetByID)
	routeV1.Put("/:id", sc.Update)
}
