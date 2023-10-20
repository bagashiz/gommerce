package app

import (
	database "github.com/bagashiz/gommerce/internal/database/mysql"
	config "github.com/bagashiz/gommerce/internal/pkg/config/viper"
	logger "github.com/bagashiz/gommerce/internal/pkg/log/zap"
)

// Run is the entrypoint of the application, dependencies are injected here
func Run() {
	log, err := logger.New()
	if err != nil {
		panic(err)
	}
	defer log.Close()

	cfgLoader, err := config.New()
	if err != nil {
		log.Fatal("failed to initialize the config provider", "error", err)
	}

	cfg, err := cfgLoader.Load()
	if err != nil {
		log.Fatal("failed to load the config", "error", err)
	}

	log.Info("succeed to load the config")
	log.Debug("config list", "config", cfg)

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

	log.Info("starting the application", "name", cfg.App.Name, "environment", cfg.App.Env)
}
