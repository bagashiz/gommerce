package main

import (
	database "github.com/bagashiz/gommerce/internal/store/mysql"
	config "github.com/bagashiz/gommerce/internal/util/config/viper"
	logger "github.com/bagashiz/gommerce/internal/util/log/zap"
)

func main() {
	log, err := logger.New()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := log.Close(); err != nil {
			log.Fatal("failed to close the logger", "error", err)
		}
	}()

	cfgProvier, err := config.New()
	if err != nil {
		log.Fatal("failed to initialize the config provider", "error", err)
	}

	cfg, err := cfgProvier.Load()
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
