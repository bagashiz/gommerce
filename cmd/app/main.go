package main

import (
	config "github.com/bagashiz/gommerce/internal/util/config/viper"
	logger "github.com/bagashiz/gommerce/internal/util/log/zap"
)

func main() {
	log, err := logger.New()
	if err != nil {
		panic(err)
	}

	cfgProvier, err := config.New()
	if err != nil {
		log.Fatal("failed to initialize the config provider", "error", err)
	}

	cfg, err := cfgProvier.Load()
	if err != nil {
		log.Fatal("failed to load the config", "error", err)
	}

	log.Info("starting the application", "name", cfg.App.Name, "environment", cfg.App.Env)
}
