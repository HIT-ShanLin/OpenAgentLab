package bootstrap

import "github.com/HIT-ShanLin/OpenAgentLab/server/pkg/config"

// Config holds network-specific configuration.
type Config struct {
	Port       int    `yaml:"port"`
	BridgeName string `yaml:"bridge_name"`
}

// LoadConfig reads configuration from file and environment.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port:       9103,
		BridgeName: "oal-br0",
	}
	if err := config.Load("configs/network", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
