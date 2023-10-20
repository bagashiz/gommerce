package viper

import (
	"os"
	"path/filepath"

	"github.com/bagashiz/gommerce/internal/pkg/config"
	"github.com/spf13/viper"
)

// Viper is a wrapper for configuration provider using viper library.
type Viper struct {
	*viper.Viper
}

// New creates a new Config instance.
func New() (config.Loader, error) {
	v := viper.New()

	// get the root directory of the application to find the config file.
	path, err := os.Executable()
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(path)

	v.AddConfigPath(dir)
	v.SetConfigFile(".env")
	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Viper{
		v,
	}, nil
}

// Load initializes the application and its dependencies configuration.
func (v *Viper) Load() (*config.Container, error) {
	app, err := v.newAppConfig()
	if err != nil {
		return nil, err
	}

	http, err := v.newHttpConfig()
	if err != nil {
		return nil, err
	}

	db, err := v.newDBConfig()
	if err != nil {
		return nil, err
	}

	return &config.Container{
		App:      app,
		Http:     http,
		Database: db,
	}, nil
}

// newAppConfig initializes the application configuration.
func (v *Viper) newAppConfig() (*config.App, error) {
	var app *config.App

	err := v.Viper.Unmarshal(&app)
	if err != nil {
		return nil, err
	}

	return app, nil
}

// newHttpConfig initializes the http server configuration.
func (v *Viper) newHttpConfig() (*config.Http, error) {
	var http *config.Http

	err := v.Viper.Unmarshal(&http)
	if err != nil {
		return nil, err
	}

	return http, nil
}

// newDBConfig initializes the database configuration.
func (v *Viper) newDBConfig() (*config.Database, error) {
	var db *config.Database

	err := v.Viper.Unmarshal(&db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
