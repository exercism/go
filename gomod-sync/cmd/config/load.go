package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load loads a configuration from a path to a JSON file.
func Load(file string) (VersionConfig, error) {
	var config VersionConfig

	f, err := os.Open(file)
	if err != nil {
		return config, fmt.Errorf("failed to open config file: %v", err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		return config, fmt.Errorf("failed to decode config file: %v", err)
	}

	return config, nil
}
