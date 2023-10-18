package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type (
	// Config is a wrapper for configuration provider using viper library.
	Config struct {
		v *viper.Viper
	}

	// Container contains the configuration for the application and its dependencies.
	Container struct {
		App      *App
		Http     *Http
		Database *Database
	}

	// App contains the configuration for the application.
	App struct {
		Name string `mapstructure:"APP_NAME"`
		Env  string `mapstructure:"APP_ENV"`
	}

	// Http contains the configuration for the http server.
	Http struct {
		Url            string `mapstructure:"HTTP_URL"`
		Port           string `mapstructure:"HTTP_PORT"`
		AllowedOrigins string `mapstructure:"HTTP_ALLOWED_ORIGINS"`
	}

	// Database contains the configuration for the database.
	Database struct {
		Conn        string `mapstructure:"DB_CONNECTION"`
		Host        string `mapstructure:"DB_HOST"`
		Port        string `mapstructure:"DB_PORT"`
		Name        string `mapstructure:"DB_NAME"`
		Username    string `mapstructure:"DB_USERNAME"`
		Password    string `mapstructure:"DB_PASSWORD"`
		MaxLifeTime int    `mapstructure:"DB_MAX_LIFE_TIME"`
		MaxOpenConn int    `mapstructure:"DB_MAX_OPEN_CONNECTIONS"`
		MaxIdleConn int    `mapstructure:"DB_MAX_IDLE_CONNECTIONS"`
	}
)

// New creates a new Config instance.
func New() (ConfigProvider, error) {
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
func (c *Config) Load() (*Container, error) {
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

	return &Container{
		app,
		http,
		db,
	}, nil
}

// newAppConfig initializes the application configuration.
func (c *Config) newAppConfig() (*App, error) {
	var app *App

	err := c.v.Unmarshal(&app)
	if err != nil {
		return nil, err
	}

	return app, nil
}

// newHttpConfig initializes the http server configuration.
func (c *Config) newHttpConfig() (*Http, error) {
	var http *Http

	err := c.v.Unmarshal(&http)
	if err != nil {
		return nil, err
	}

	return http, nil
}

// newDBConfig initializes the database configuration.
func (c *Config) newDBConfig() (*Database, error) {
	var db *Database

	err := c.v.Unmarshal(&db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
