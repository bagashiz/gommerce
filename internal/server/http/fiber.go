package http

import (
	"github.com/bagashiz/gommerce/internal/pkg/config"
	"github.com/bagashiz/gommerce/internal/pkg/log"
	"github.com/bagashiz/gommerce/internal/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Fiber is a wrapper for Fiber router
type Fiber struct {
	App    *fiber.App
	cfg    *config.Http
	logger log.Logger
}

// New creates a new Fiber router instance
func New(cfg *config.Http, logger log.Logger) server.Http {
	app := fiber.New()
	return &Fiber{
		app,
		cfg,
		logger,
	}
}

// InitRoutes initializes the routes for the Fiber router
func (f *Fiber) InitRoutes() {
	logger := f.customLogger()

	cors := cors.New(cors.Config{
		AllowOrigins: f.cfg.AllowedOrigins,
	})

	f.App.Use(logger, cors)

	api := f.App.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is version 1 of the API")
	})
}

// Start starts the Fiber router
func (f *Fiber) Start() error {
	listenAddr := f.cfg.Url + ":" + f.cfg.Port
	return f.App.Listen(listenAddr)
}

// customLogger is a custom logger for Fiber
func (f *Fiber) customLogger() fiber.Handler {
	format := `${status} | ${method} | ${path} | ${protocol} | ${ip} | ${latency} | ${ua}`
	return logger.New(logger.Config{
		Format:        format,
		DisableColors: true,
		Done: func(c *fiber.Ctx, logString []byte) {
			f.logger.Info("request", "log", string(logString))
		},
	})
}
