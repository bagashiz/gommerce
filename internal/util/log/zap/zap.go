package zap

import (
	"os"

	"github.com/bagashiz/gommerce/internal/util/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a wrapper for the log interface using zap library.
type Logger struct {
	zap *zap.SugaredLogger
}

// filename is the path to the file where the logs will be written.
const filename = "app.log"

// New creates a new Log instance.
func New() (log.LogProvider, error) {
	var (
		cores   []zapcore.Core
		logFile *os.File
		err     error
	)

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	if os.Getenv("APP_ENV") == "production" {
		// Open the log file for production
		logFile, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}

		// Create a file core
		fileCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			zapcore.AddSync(logFile),
			zapcore.InfoLevel,
		)
		cores = append(cores, fileCore)
	}

	// Create a console core
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)
	cores = append(cores, consoleCore)

	// Create the logger with additional context information (caller, stack trace)
	logger := zap.New(
		zapcore.NewTee(cores...),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	).Sugar()

	defer func() {
		_ = logger.Sync()
	}()

	return &Logger{
		logger,
	}, nil
}

// Debug logs a debug message.
func (l *Logger) Debug(format string, args ...any) {
	l.zap.Debugw(format, args...)
}

// Info logs an info message.
func (l *Logger) Info(format string, args ...any) {
	l.zap.Infow(format, args...)
}

// Warn logs a warning message.
func (l *Logger) Warn(format string, args ...any) {
	l.zap.Warnw(format, args...)
}

// Error logs an error message.
func (l *Logger) Error(format string, args ...any) {
	l.zap.Errorw(format, args...)
}

// Fatal logs a fatal message.
func (l *Logger) Fatal(format string, args ...any) {
	l.zap.Fatalw(format, args...)
}

// Panic logs a panic message.
func (l *Logger) Panic(format string, args ...any) {
	l.zap.Panicw(format, args...)
}
