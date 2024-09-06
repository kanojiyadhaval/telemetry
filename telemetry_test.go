package telemetry

import (
	"testing"

	"github.com/kanojiyadhaval/telemetry/drivers"
)

func TestTelemetryLog(t *testing.T) {
	logger := drivers.NewCLI()
	telemetry := NewTelemetry(logger)
	attributes := map[string]string{"origin": "http", "customerId": "123"}
	telemetry.Log(Info, "Test log", attributes)
}
