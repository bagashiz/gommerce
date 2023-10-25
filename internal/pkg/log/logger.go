package log

// Logger is an interface that defines the methods for the logger.
type Logger interface {
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
	// Close flushes any buffered log entries.
	Close() error
}

// New initializes the logger.
func New() (Logger, error) {
	return newZap()
}
