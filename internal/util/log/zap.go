package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a wrapper for the log interface using zap library.
type Logger struct {
	zap *zap.SugaredLogger
}

// logFile is the path to the file where the logs will be written.
const logFile = "app.log"

// New creates a new Log instance.
func New() (LogProvider, error) {
	// configure the time format
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	// create file and console encoders
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	// open the log file
	logFile, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	// create writers for file and console
	fileWriter := zapcore.AddSync(logFile)
	consoleWriter := zapcore.AddSync(os.Stdout)

	// set the log level
	defaultLogLevel := zapcore.DebugLevel
	if os.Getenv("APP_ENV") == "production" {
		defaultLogLevel = zapcore.InfoLevel
	}

	// create cores for writing to the file and console
	fileCore := zapcore.NewCore(fileEncoder, fileWriter, defaultLogLevel)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleWriter, defaultLogLevel)

	// combine cores
	core := zapcore.NewTee(fileCore, consoleCore)

	// create the logger with additional context information (caller, stack trace)
	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	).Sugar()

	defer logger.Sync()

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
