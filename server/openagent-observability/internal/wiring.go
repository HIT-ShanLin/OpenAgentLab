package observability

import (
	"fmt"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-observability/internal/adapter"
	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-observability/internal/service"
	"github.com/HIT-ShanLin/OpenAgentLab/server/pkg/config"
	"github.com/HIT-ShanLin/OpenAgentLab/server/pkg/logger"
)

// Config holds observability-specific configuration.
type Config struct {
	Port             int `yaml:"port"`
	MetricsPort      int `yaml:"metrics_port"`
}

// LoadConfig reads configuration.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port:        9104,
		MetricsPort: 9200,
	}
	if err := config.Load("configs/observability", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// App is the assembled observability application.
type App struct {
	Service *service.MetricsService
	Logger  *logger.Logger
	Config  *Config
}

// Wire assembles all dependencies.
func Wire(cfg *Config) (*App, error) {
	log, err := logger.New()
	if err != nil {
		return nil, fmt.Errorf("wiring logger: %w", err)
	}

	collector := adapter.NewPrometheusCollector(cfg.MetricsPort)
	svc := service.New(collector)

	log.Info("observability service wired successfully", "port", cfg.Port, "metrics_port", cfg.MetricsPort)

	return &App{Service: svc, Logger: log, Config: cfg}, nil
}

// Run starts the application.
func (a *App) Run() error {
	a.Logger.Info("openagent-observability starting", "port", a.Config.Port)
	return nil
}
