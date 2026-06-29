package bootstrap

import (
	"fmt"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-network/internal/adapter"
	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-network/internal/repo"
	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-network/internal/service"
	"github.com/HIT-ShanLin/OpenAgentLab/server/pkg/logger"
)

// App is the assembled network application.
type App struct {
	Service *service.NetworkService
	Logger  *logger.Logger
	Config  *Config
}

// Wire assembles all dependencies.
func Wire(cfg *Config) (*App, error) {
	log, err := logger.New()
	if err != nil {
		return nil, fmt.Errorf("wiring logger: %w", err)
	}

	memRepo := repo.NewMemoryRepository()
	driver := adapter.NewLinuxBridge(cfg.BridgeName)
	svc := service.New(memRepo, driver)

	log.Info("network service wired successfully", "port", cfg.Port, "bridge", cfg.BridgeName)

	return &App{Service: svc, Logger: log, Config: cfg}, nil
}

// Run starts the application.
func (a *App) Run() error {
	a.Logger.Info("openagent-network starting", "port", a.Config.Port)
	return nil
}
