package telemetry

import (
	"time"
)

// LogLevel represents the severity of the log
type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

func (l LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARNING", "ERROR"}[l]
}

// LogMessage represents a log entry with meta-information
type LogMessage struct {
	Timestamp  time.Time
	Level      LogLevel
	Message    string
	Attributes map[string]string
}

// Logger is an interface that defines the behavior of a logging driver
type Logger interface {
	Log(msg LogMessage)
}

// Telemetry is the main struct holding the logger and configuration
type Telemetry struct {
	logger Logger
}

// NewTelemetry creates a new Telemetry instance
func NewTelemetry(logger Logger) *Telemetry {
	return &Telemetry{logger: logger}
}

// Log logs a message with the specified log level and attributes
func (t *Telemetry) Log(level LogLevel, message string, attributes map[string]string) {
	logMsg := LogMessage{
		Timestamp:  time.Now(),
		Level:      level,
		Message:    message,
		Attributes: attributes,
	}
	t.logger.Log(logMsg)
}

// TransactionLog starts a transaction-based logging session
func (t *Telemetry) TransactionLog(transactionID string, attributes map[string]string) {
	attributes["transaction_id"] = transactionID
	t.Log(Info, "Transaction started", attributes)
}
