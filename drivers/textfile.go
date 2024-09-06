package drivers

import (
	"fmt"
	"os"
	"telemetry"
	"time"
)

type TextFile struct {
	filePath string
}

func NewTextFile(filePath string) *TextFile {
	return &TextFile{filePath: filePath}
}

func (t *TextFile) Log(msg telemetry.LogMessage) {
	file, err := os.OpenFile(t.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	logEntry := fmt.Sprintf("[%s] [%s] %s - Attributes: %v\n", msg.Timestamp.Format(time.RFC3339), msg.Level.String(), msg.Message, msg.Attributes)
	file.WriteString(logEntry)
}
