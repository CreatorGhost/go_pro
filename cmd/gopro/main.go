package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"yourprojectname/handlers"
)

type ServerConfig struct {
	Port string `json:"port"`
}

func LoadServerConfig() (ServerConfig, error) {
	var config ServerConfig
	file, err := os.Open("config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	if config.Port == "" {
		config.Port = "8000"
	}

	return config, nil
}

func main() {
	config, err := LoadServerConfig()
	if err != nil {
		log.Fatalf("Error loading server configuration: %v", err)
	}

	http.HandleFunc("/process-single", handlers.ProcessSingleHandler)
	http.HandleFunc("/process-concurrent", handlers.ProcessConcurrentHandler)

	log.Printf("Starting server on port %s", config.Port)
	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}