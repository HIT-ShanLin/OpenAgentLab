// Package config provides YAML configuration loading with environment variable overrides.
package config

import (
	"fmt"
	"os"
	"strings"
)

// Load reads a YAML config file and populates the given struct.
// It searches for <path>.yaml in the current directory and config paths.
// This is a minimal stub that sets defaults; a full implementation would parse YAML.
func Load(path string, cfg interface{}) error {
	// Search paths: <path>.yaml, server/<path>.yaml
	candidates := []string{
		path + ".yaml",
		"server/" + path + ".yaml",
	}

	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			// TODO: parse YAML with gopkg.in/yaml.v3
			_ = p
			return nil
		}
	}

	// No config file found — use defaults (caller sets them before calling Load)
	if os.Getenv("OPENAGENT_CONFIG_STRICT") != "" {
		return fmt.Errorf("config: no config file found for %s (searched: %s)", path, strings.Join(candidates, ", "))
	}
	return nil
}
