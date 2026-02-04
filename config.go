package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	ApiKey string `json:"api_key"`
}

func configPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "syniq", "config.json")
}

func saveConfig(cfg Config) {
	path := configPath()
	os.MkdirAll(filepath.Dir(path), 0700)

	data, _ := json.Marshal(cfg)
	os.WriteFile(path, data, 0600)
}

func loadConfig() (Config, error) {
	var cfg Config
	data, err := os.ReadFile(configPath())
	if err != nil {
		return cfg, fmt.Errorf("not connected")
	}
	json.Unmarshal(data, &cfg)
	return cfg, nil
}
