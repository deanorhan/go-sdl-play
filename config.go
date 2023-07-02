package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Display struct {
	Width  int32 `yaml:"width"`
	Height int32 `yaml:"height"`
}

type Config struct {
	Display Display `yaml:"display"`
	Debug   bool    `yaml:"debug"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
