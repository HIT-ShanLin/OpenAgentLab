package bootstrap

import (
	"fmt"

	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/sandbox/adapter"
	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/sandbox/repo"
	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/sandbox/service"
	"github.com/HIT-ShanLin/OpenAgentLab/server/pkg/logger"
)

// App is the assembled sandbox application.
type App struct {
	Service *service.SandboxService
	Logger  *logger.Logger
	Config  *Config
}

// Wire assembles all dependencies (manual DI).
func Wire(cfg *Config) (*App, error) {
	log, err := logger.New()
	if err != nil {
		return nil, fmt.Errorf("wiring logger: %w", err)
	}

	memRepo := repo.NewMemoryRepository()
	runner := adapter.NewNamespaceRunner()
	svc := service.New(memRepo, runner)

	log.Info("sandbox service wired successfully", "port", cfg.Port)

	return &App{
		Service: svc,
		Logger:  log,
		Config:  cfg,
	}, nil
}

// Run starts the application (stub — will be replaced with gRPC/HTTP server).
func (a *App) Run() error {
	a.Logger.Info("openagent-sandbox starting", "port", a.Config.Port)
	// TODO: start gRPC/HTTP server
	return nil
}
