package main

import (
	"github.com/bagashiz/gommerce/internal/util/config"
	"github.com/bagashiz/gommerce/internal/util/log"
)

func main() {
	log, err := log.New()
	if err != nil {
		panic(err)
	}

	cp, err := config.New()
	if err != nil {
		log.Fatal("failed to initialize the config provider", "error", err)
	}

	cfg, err := cp.Load()
	if err != nil {
		log.Fatal("failed to load the config", "error", err)
	}

	log.Info("starting the application", "name", cfg.App.Name, "environment", cfg.App.Env)
}
