package store

// DB is an interface that defines the methods for the database.
type DB interface {
	// Migrate runs the auto migration for the database.
	Migrate() error
	// Close closes the database connection.
	Close() error
}
