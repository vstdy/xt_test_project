package api

import (
	"fmt"
	"time"
)

// Config keeps rest params.
type Config struct {
	Timeout    time.Duration `mapstructure:"timeout"`
	RunAddress string        `mapstructure:"run_address"`
}

// Validate performs a basic validation.
func (config Config) Validate() error {
	if config.RunAddress == "" {
		return fmt.Errorf("%s field: empty", "server_address")
	}

	return nil
}

// NewDefaultConfig builds a Config with default values.
func NewDefaultConfig() Config {
	return Config{
		Timeout:    5 * time.Second,
		RunAddress: "0.0.0.0:8080",
	}
}
