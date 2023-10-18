package log

// LogProvider is an interface that defines the methods for the logger.
type LogProvider interface {
	// Debug logs a debug message.
	Debug(format string, args ...any)
	// Info logs an info message.
	Info(format string, args ...any)
	// Warn logs a warning message.
	Warn(format string, args ...any)
	// Error logs an error message.
	Error(format string, args ...any)
	// Fatal logs a fatal message.
	Fatal(format string, args ...any)
	// Panic logs a panic message.
	Panic(format string, args ...any)
}
