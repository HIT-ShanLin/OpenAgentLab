package bootstrap

import (
	"fmt"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-runtime/internal/adapter"
	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-runtime/internal/repo"
	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-runtime/internal/service"
	"github.com/HIT-ShanLin/OpenAgentLab/server/pkg/logger"
)

// App is the assembled runtime application.
type App struct {
	Service *service.RuntimeService
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
	starter := adapter.NewContainerStarter()
	svc := service.New(memRepo, starter)

	log.Info("runtime service wired successfully", "port", cfg.Port)

	return &App{Service: svc, Logger: log, Config: cfg}, nil
}

// Run starts the application.
func (a *App) Run() error {
	a.Logger.Info("openagent-runtime starting", "port", a.Config.Port)
	return nil
}
