package database

import (
	"github.com/bagashiz/gommerce/internal/pkg/config"
	"github.com/bagashiz/gommerce/internal/pkg/helper"
	"gorm.io/gorm"
)

// DB is an interface that defines the methods for the GORM database connection.
type DB interface {
	// Migrate runs the auto migration for the database.
	Migrate() error
	// DB returns the underlying gorm.DB connection.
	DB() *gorm.DB
	// Close closes the database connection.
	Close() error
}

// New initializes the database connection based on the configuration.
func New(cfg *config.Database) (DB, error) {
	switch cfg.Conn {
	case "mysql":
		return newMysql(cfg)
		// TODO: Add support for other databases
	}

	return nil, helper.ErrUnsupportedDriver
}
