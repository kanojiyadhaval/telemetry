package telemetry

import (
	"telemetry/drivers"
	"testing"
)

func TestTelemetryLog(t *testing.T) {
	logger := drivers.NewCLI()
	telemetry := NewTelemetry(logger)
	attributes := map[string]string{"origin": "http", "customerId": "123"}
	telemetry.Log(Info, "Test log", attributes)
}
