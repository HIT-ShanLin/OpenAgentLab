// Package logger provides structured logging via slog.
// All OpenAgentLab services use this package for consistent log formatting.
package logger

import (
	"log/slog"
	"os"
)

// Logger wraps slog.Logger with convenience methods.
type Logger struct {
	*slog.Logger
}

// New creates a new structured logger writing to stdout with JSON format.
func New() (*Logger, error) {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	return &Logger{Logger: slog.New(handler)}, nil
}

// NewWithLevel creates a logger with a custom log level.
func NewWithLevel(level slog.Level) *Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})
	return &Logger{Logger: slog.New(handler)}
}

// Info logs at info level with structured fields.
func (l *Logger) Info(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

// Error logs at error level with structured fields.
func (l *Logger) Error(msg string, args ...any) {
	l.Logger.Error(msg, args...)
}

// Debug logs at debug level with structured fields.
func (l *Logger) Debug(msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}

// Warn logs at warn level with structured fields.
func (l *Logger) Warn(msg string, args ...any) {
	l.Logger.Warn(msg, args...)
}
