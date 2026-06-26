package bootstrap

import "github.com/HIT-ShanLin/OpenAgentLab/server/pkg/config"

// Config holds runtime-specific configuration.
type Config struct {
	Port         int    `yaml:"port"`
	ContainerBin string `yaml:"container_bin"`
}

// LoadConfig reads configuration from file and environment.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port:         9101,
		ContainerBin: "/usr/bin/runc",
	}
	if err := config.Load("configs/runtime", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
