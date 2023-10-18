package viper

import (
	"os"
	"path/filepath"

	"github.com/bagashiz/gommerce/internal/util/config"
	"github.com/spf13/viper"
)

// Config is a wrapper for configuration provider using viper library.
type Config struct {
	v *viper.Viper
}

// New creates a new Config instance.
func New() (config.ConfigProvider, error) {
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

	return &Config{
		v,
	}, nil
}

// Load initializes the application and its dependencies configuration.
func (c *Config) Load() (*config.Container, error) {
	app, err := c.newAppConfig()
	if err != nil {
		return nil, err
	}

	http, err := c.newHttpConfig()
	if err != nil {
		return nil, err
	}

	db, err := c.newDBConfig()
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
func (c *Config) newAppConfig() (*config.App, error) {
	var app *config.App

	err := c.v.Unmarshal(&app)
	if err != nil {
		return nil, err
	}

	return app, nil
}

// newHttpConfig initializes the http server configuration.
func (c *Config) newHttpConfig() (*config.Http, error) {
	var http *config.Http

	err := c.v.Unmarshal(&http)
	if err != nil {
		return nil, err
	}

	return http, nil
}

// newDBConfig initializes the database configuration.
func (c *Config) newDBConfig() (*config.Database, error) {
	var db *config.Database

	err := c.v.Unmarshal(&db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
