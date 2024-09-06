package drivers

import (
	"fmt"
	"time"

	"github.com/kanojiyadhaval/telemetry"
)

// CLI is a logger implementation that logs to the console
type CLI struct{}

func NewCLI() *CLI {
	return &CLI{}
}

func (c *CLI) Log(msg telemetry.LogMessage) {
	fmt.Printf("[%s] [%s] %s - Attributes: %v\n", msg.Timestamp.Format(time.RFC3339), msg.Level.String(), msg.Message, msg.Attributes)
}
