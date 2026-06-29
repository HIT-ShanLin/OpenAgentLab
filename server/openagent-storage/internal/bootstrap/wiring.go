package bootstrap

import (
	"fmt"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-storage/internal/adapter"
	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-storage/internal/repo"
	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-storage/internal/service"
	"github.com/HIT-ShanLin/OpenAgentLab/server/pkg/logger"
)

// App is the assembled storage application.
type App struct {
	Service *service.VolumeService
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
	provisioner := adapter.NewOverlayFSProvisioner(cfg.BasePath)
	svc := service.New(memRepo, provisioner)

	log.Info("storage service wired successfully", "port", cfg.Port, "base_path", cfg.BasePath)

	return &App{Service: svc, Logger: log, Config: cfg}, nil
}

// Run starts the application.
func (a *App) Run() error {
	a.Logger.Info("openagent-storage starting", "port", a.Config.Port)
	return nil
}
