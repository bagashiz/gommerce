package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Viper is a wrapper for configuration provider using viper library.
type Viper struct {
	*viper.Viper
}

// newViper creates a new Config instance.
func newViper() (Config, error) {
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
func (v *Viper) Load() (*Container, error) {
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

	token, err := v.newTokenConfig()
	if err != nil {
		return nil, err
	}

	return &Container{
		App:      app,
		Http:     http,
		Database: db,
		Token:    token,
	}, nil
}

// newAppConfig initializes the application configuration.
func (v *Viper) newAppConfig() (*App, error) {
	var app App

	err := v.Viper.Unmarshal(&app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

// newHttpConfig initializes the http server configuration.
func (v *Viper) newHttpConfig() (*Http, error) {
	var http Http

	err := v.Viper.Unmarshal(&http)
	if err != nil {
		return nil, err
	}

	return &http, nil
}

// newDBConfig initializes the database configuration.
func (v *Viper) newDBConfig() (*Database, error) {
	var db Database

	err := v.Viper.Unmarshal(&db)
	if err != nil {
		return nil, err
	}

	return &db, nil
}

// newTokenConfig initializes the token configuration.
func (v *Viper) newTokenConfig() (*Token, error) {
	var token Token

	err := v.Viper.Unmarshal(&token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
