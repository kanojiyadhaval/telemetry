package telemetry

import "time"

// Log represents a structured log entry with a timestamp, log level, message, and attributes.
type Log struct {
	Timestamp  time.Time         `json:"timestamp"`
	Level      LogLevel          `json:"level"`
	Message    string            `json:"message"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

// NewLog creates a new log entry
func NewLog(level LogLevel, message string, attributes map[string]string) Log {
	return Log{
		Timestamp:  time.Now(),
		Level:      level,
		Message:    message,
		Attributes: attributes,
	}
}
