package bootstrap

import "github.com/HIT-ShanLin/OpenAgentLab/server/pkg/config"

// Config holds control-plane-specific configuration.
type Config struct {
	Port      int `yaml:"port"`
	QueueSize int `yaml:"queue_size"`
}

// LoadConfig reads configuration from file and environment.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port:      9105,
		QueueSize: 1000,
	}
	if err := config.Load("configs/controlplane", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
