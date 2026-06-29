package bootstrap

import (
	"github.com/HIT-ShanLin/OpenAgentLab/server/pkg/config"
)

// Config holds sandbox-specific configuration.
type Config struct {
	Port       int    `yaml:"port"`
	RuntimeDir string `yaml:"runtime_dir"`
}

// LoadConfig reads configuration from file and environment.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port:       9100,
		RuntimeDir: "/var/run/openagent/sandbox",
	}
	if err := config.Load("configs/sandbox", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
