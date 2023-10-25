package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap is a wrapper for the log interface using zap library.
type Zap struct {
	*zap.SugaredLogger
}

// filename is the path to the file where the logs will be written.
const filename = "app.log"

// newZap creates a new Log instance.
func newZap() (Logger, error) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	// Open the log file
	logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	// Create a file core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(logFile),
		zap.InfoLevel,
	)

	// Create the logger with additional context information (caller, stack trace)
	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	).Sugar()

	return &Zap{
		logger,
	}, nil
}

// Debug logs a debug message.
func (z *Zap) Debug(format string, args ...any) {
	z.Debugw(format, args...)
}

// Info logs an info message.
func (z *Zap) Info(format string, args ...any) {
	z.Infow(format, args...)
}

// Warn logs a warning message.
func (z *Zap) Warn(format string, args ...any) {
	z.Warnw(format, args...)
}

// Error logs an error message.
func (z *Zap) Error(format string, args ...any) {
	z.Errorw(format, args...)
}

// Fatal logs a fatal message.
func (z *Zap) Fatal(format string, args ...any) {
	z.Fatalw(format, args...)
}

// Panic logs a panic message.
func (z *Zap) Panic(format string, args ...any) {
	z.Panicw(format, args...)
}

// Close flushes any buffered log entries.
func (z *Zap) Close() error {
	if err := z.Sync(); err != nil {
		return err
	}

	return nil
}
