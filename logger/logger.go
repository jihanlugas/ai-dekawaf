package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	// Set output to stdout
	Log.SetOutput(os.Stdout)

	// Set JSON formatter for structured logging
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	})

	// Set log level based on environment variable
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	switch logLevel {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "info":
		Log.SetLevel(logrus.InfoLevel)
	case "warn", "warning":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	default:
		// Default to Info level in production
		if os.Getenv("ENV") == "local" {
			Log.SetLevel(logrus.DebugLevel)
		} else {
			Log.SetLevel(logrus.InfoLevel)
		}
	}
}

// Debug logs a debug message with optional fields
func Debug(msg string, fields ...logrus.Fields) {
	if len(fields) > 0 {
		Log.WithFields(fields[0]).Debug(msg)
	} else {
		Log.Debug(msg)
	}
}

// Info logs an info message with optional fields
func Info(msg string, fields ...logrus.Fields) {
	if len(fields) > 0 {
		Log.WithFields(fields[0]).Info(msg)
	} else {
		Log.Info(msg)
	}
	fmt.Println("")
}

// Warn logs a warning message with optional fields
func Warn(msg string, fields ...logrus.Fields) {
	if len(fields) > 0 {
		Log.WithFields(fields[0]).Warn(msg)
	} else {
		Log.Warn(msg)
	}
}

// Error logs an error message with optional fields
func Error(msg string, err error, fields ...logrus.Fields) {
	logFields := logrus.Fields{}
	if len(fields) > 0 {
		logFields = fields[0]
	}
	if err != nil {
		logFields["error"] = err.Error()
	}
	Log.WithFields(logFields).Error(msg)
	fmt.Println("")
}

// Fatal logs a fatal message and exits
func Fatal(msg string, err error, fields ...logrus.Fields) {
	logFields := logrus.Fields{}
	if len(fields) > 0 {
		logFields = fields[0]
	}
	if err != nil {
		logFields["error"] = err.Error()
	}
	Log.WithFields(logFields).Fatal(msg)
}

// WithFields creates a logger with predefined fields
func WithFields(fields logrus.Fields) *logrus.Entry {
	return Log.WithFields(fields)
}
