package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DriverType string `json:"driver_type"`
	FilePath   string `json:"file_path"`
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
