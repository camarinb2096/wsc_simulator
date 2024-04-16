package logger

import (
	"log"
	"os"
)

// Log interface outlines the logging functions
type Log interface {
	Debug(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
}

// Logger struct embeds a *log.Logger from the standard library
type Logger struct {
	logger *log.Logger
}

// NewLogger initializes a new Logger instance
func NewLogger() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "WSC-Simulator", log.LstdFlags),
	}
}

// Debug logs a debug message
func (l *Logger) Debug(message string, args ...interface{}) {
	l.logger.Printf("DEBUG: "+message, args...)
}

// Error logs an error message
func (l *Logger) Error(message string, args ...interface{}) {
	l.logger.Printf("ERROR: "+message, args...)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(message string, args ...interface{}) {
	l.logger.Fatalf("FATAL: "+message, args...)
}

// Info logs an informational message
func (l *Logger) Info(message string, args ...interface{}) {
	l.logger.Printf("INFO: "+message, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(message string, args ...interface{}) {
	l.logger.Printf("WARN: "+message, args...)
}
