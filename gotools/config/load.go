package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Load loads a configuration from a path to a JSON file.
func Load(file string) (Config, error) {
	var config Config

	jsonData, err := os.ReadFile(file)
	if err != nil {
		return config, fmt.Errorf("failed to open config file: %v", err)
	}

	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		return config, fmt.Errorf("failed to decode config file: %v", err)
	}
	return config, nil
}

// Return the default go.mod version from a config.
func DefaultVersion(file string) (string, error) {
	c, err := Load(file)
	if err != nil {
		return "", err
	}
	if c.GoModVersion.Default == "" {
		return "", errors.New("Did not find a default go.mod version in the config")
	}
	return c.GoModVersion.Default, nil
}
