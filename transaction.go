package telemetry

import "fmt"

// Transaction represents a log entry for a transaction.
type Transaction struct {
	ID         string            // Unique ID for the transaction
	Attributes map[string]string // Additional metadata for the transaction
	telemetry  *Telemetry        // Reference to the telemetry instance
}

// NewTransaction creates a new transaction log.
func NewTransaction(telemetry *Telemetry, transactionID string, attributes map[string]string) *Transaction {
	return &Transaction{
		ID:         transactionID,
		Attributes: attributes,
		telemetry:  telemetry,
	}
}

// Start logs the start of a transaction.
func (t *Transaction) Start() {
	message := fmt.Sprintf("Transaction %s started", t.ID)
	t.telemetry.Log(Info, message, t.Attributes)
}

// Log logs a message as part of this transaction.
func (t *Transaction) Log(level LogLevel, message string, attributes map[string]string) {
	// Merge transaction-level attributes with log-specific attributes
	allAttributes := mergeAttributes(t.Attributes, attributes)
	t.telemetry.Log(level, message, allAttributes)
}

// End logs the end of a transaction.
func (t *Transaction) End() {
	message := fmt.Sprintf("Transaction %s ended", t.ID)
	t.telemetry.Log(Info, message, t.Attributes)
}

// mergeAttributes merges two maps of attributes (transaction-level and log-specific attributes).
func mergeAttributes(transAttrs, logAttrs map[string]string) map[string]string {
	merged := make(map[string]string)

	for k, v := range transAttrs {
		merged[k] = v
	}
	for k, v := range logAttrs {
		merged[k] = v
	}

	return merged
}
