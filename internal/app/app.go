package app

import (
	"github.com/bagashiz/gommerce/internal/app/address"
	"github.com/bagashiz/gommerce/internal/app/auth"
	"github.com/bagashiz/gommerce/internal/app/category"
	"github.com/bagashiz/gommerce/internal/app/city"
	"github.com/bagashiz/gommerce/internal/app/province"
	"github.com/bagashiz/gommerce/internal/app/shop"
	"github.com/bagashiz/gommerce/internal/app/user"
	"github.com/bagashiz/gommerce/internal/pkg/config"
	"github.com/bagashiz/gommerce/internal/pkg/database"
	"github.com/bagashiz/gommerce/internal/pkg/log"
	"github.com/bagashiz/gommerce/internal/pkg/server/http"
	"github.com/bagashiz/gommerce/internal/pkg/token"
)

// Run is the entrypoint of the application, dependencies are injected here
func Run() {
	log, err := log.New()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := log.Close(); err != nil {
			log.Panic("failed to close the logger", "error", err)
		}
	}()

	cfgLoader, err := config.New()
	if err != nil {
		log.Fatal("failed to initialize the config provider", "error", err)
	}

	cfg, err := cfgLoader.Load()
	if err != nil {
		log.Fatal("failed to load the config", "error", err)
	}

	log.Info("succeed to load the config")

	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal("failed to initialize the database", "error", err)
	}
	defer db.Close()

	log.Info("succeed to initialize the database", "connection", cfg.Database.Conn)

	if err := db.Migrate(); err != nil {
		log.Fatal("failed to migrate the database", "error", err)
	}

	log.Info("succeed to migrate the database")

	token, err := token.New(cfg.Token)
	if err != nil {
		log.Fatal("failed to initialize the token provider", "error", err)
	}

	log.Info("succeed to initialize the token provider", "type", cfg.Token.Type)

	server := http.New(cfg.Http, log)

	// Dependency injection
	province.New(server)
	city.New(server)
	auth.New(db, server, token)
	category.New(db, server)
	address.New(db, server)
	user.New(db, server)
	shop.New(db, server)

	log.Info("starting the application", "name", cfg.App.Name, "environment", cfg.App.Env)

	if err := server.Start(); err != nil {
		log.Fatal("failed to start the server", "error", err)
	}
}
