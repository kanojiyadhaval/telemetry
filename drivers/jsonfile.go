package drivers

import (
	"encoding/json"
	"fmt"
	"os"
	"telemetry"
)

type JSONFile struct {
	filePath string
}

func NewJSONFile(filePath string) *JSONFile {
	return &JSONFile{filePath: filePath}
}

func (j *JSONFile) Log(msg telemetry.LogMessage) {
	file, err := os.OpenFile(j.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	logEntry, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error marshalling log:", err)
		return
	}

	file.Write(logEntry)
	file.Write([]byte("\n"))
}
