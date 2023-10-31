package v1

import "github.com/gofiber/fiber/v2"

// InitRoutes registers all routes for version 1
func (ac *AuthControllerV1) InitRoutes(routeV1 fiber.Router) {
	routeV1.Post("/register", ac.Register)
	routeV1.Post("/login", ac.Login)
}
