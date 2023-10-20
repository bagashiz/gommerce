package store

import "gorm.io/gorm"

// DB is an interface that defines the methods for the database.
type DB interface {
	// Migrate runs the auto migration for the database.
	Migrate() error
	// DB returns the underlying gorm.DB connection.
	DB() *gorm.DB
	// Close closes the database connection.
	Close() error
}
