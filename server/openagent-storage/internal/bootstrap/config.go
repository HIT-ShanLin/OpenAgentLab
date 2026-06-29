package bootstrap

import "github.com/HIT-ShanLin/OpenAgentLab/server/pkg/config"

// Config holds storage-specific configuration.
type Config struct {
	Port     int    `yaml:"port"`
	BasePath string `yaml:"base_path"`
}

// LoadConfig reads configuration from file and environment.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port:     9102,
		BasePath: "/var/lib/openagent/storage",
	}
	if err := config.Load("configs/storage", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
