package config

import "time"

// Config is an interface that defines the methods for the configuration loader.
type Config interface {
	// Load initializes the application and its dependencies configuration.
	Load() (*Container, error)
}

type (
	// Container contains the configuration for the application and its dependencies.
	Container struct {
		App      *App
		Http     *Http
		Database *Database
		Token    *Token
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

	// Token contains the configuration for the token.
	Token struct {
		Type         string        `mapstructure:"TOKEN_TYPE"`
		SymmetricKey string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
		Duration     time.Duration `mapstructure:"TOKEN_DURATION"`
	}
)

// New initializes the configuration loader.
func New() (Config, error) {
	return newViper()
}
