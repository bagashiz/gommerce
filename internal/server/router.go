package server

// Http is an interface that defines the methods for the HTTP server.
type Http interface {
	// InitRoutes initializes the routes for the server.
	InitRoutes()
	// Start starts the server.
	Start() error
}
