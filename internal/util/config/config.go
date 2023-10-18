package config

// ConfigProvider is an interface that defines the methods for the configuration provider.
type ConfigProvider interface {
	// Load initializes the application and its dependencies configuration.
	Load() (*Container, error)
}
