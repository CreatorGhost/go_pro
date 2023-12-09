package utils

import (
	"encoding/json"
	"os"
)

// ServerConfig holds the configuration for the server
type ServerConfig struct {
	Port string `json:"port"`
}

// LoadServerConfig loads the server configuration from a JSON file
func LoadServerConfig() (ServerConfig, error) {
	var config ServerConfig

	// Open the configuration file
	file, err := os.Open("config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()

	// Decode the JSON configuration
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, err
	}

	// Set default port if not specified
	if config.Port == "" {
		config.Port = "8000"
	}

	return config, nil
}
