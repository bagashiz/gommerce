package http

import (
	"github.com/bagashiz/gommerce/internal/pkg/config"
	"github.com/bagashiz/gommerce/internal/pkg/log"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Http is a wrapper for Fiber router
type Http struct {
	App      *fiber.App
	Validate *validator.Validate
	Logger   log.Logger
	cfg      *config.Http
}

// New creates a new Fiber router instance
func New(cfg *config.Http, logger log.Logger) *Http {
	app := fiber.New()

	fiberLogger := customLogger(logger)

	cors := cors.New(cors.Config{
		AllowOrigins: cfg.AllowedOrigins,
	})

	app.Use(fiberLogger, cors)

	validate := validator.New()

	return &Http{
		app,
		validate,
		logger,
		cfg,
	}
}

// Start starts the Http server
func (h *Http) Start() error {
	listenAddr := h.cfg.Url + ":" + h.cfg.Port
	return h.App.Listen(listenAddr)
}

// customLogger is a custom logger for Fiber
func customLogger(l log.Logger) fiber.Handler {
	format := `${status} | ${method} | ${path} | ${protocol} | ${ip} | ${latency} | ${ua}`
	return logger.New(logger.Config{
		Format:        format,
		DisableColors: true,
		Done: func(c *fiber.Ctx, logString []byte) {
			l.Info("request", "log", string(logString))
		},
	})
}
