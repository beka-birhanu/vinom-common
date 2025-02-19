// Package i defines various general interfaces used in Vinom.
package i

// Logger defines a simple logging interface for structured logging at different levels.
type Logger interface {
	// Error logs an error message.
	Error(message string)

	// Info logs an informational message.
	Info(message string)

	// Warning logs a warning message.
	Warning(message string)
}
