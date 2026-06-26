package bootstrap

import (
	"fmt"

	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/controlplane/adapter"
	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/controlplane/repo"
	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/controlplane/service"
	"github.com/HIT-ShanLin/OpenAgentLab/server/pkg/logger"
)

// App is the assembled control-plane application.
type App struct {
	Service *service.ControlPlaneService
	Logger  *logger.Logger
	Config  *Config
}

// Wire assembles all dependencies.
func Wire(cfg *Config) (*App, error) {
	log, err := logger.New()
	if err != nil {
		return nil, fmt.Errorf("wiring logger: %w", err)
	}

	taskRepo := repo.NewMemoryTaskRepository()
	queueRepo := repo.NewMemoryQueueRepository()
	scheduler := adapter.NewSimpleScheduler()
	svc := service.New(taskRepo, queueRepo, scheduler)

	log.Info("control-plane service wired successfully", "port", cfg.Port, "queue_size", cfg.QueueSize)

	return &App{Service: svc, Logger: log, Config: cfg}, nil
}

// Run starts the application.
func (a *App) Run() error {
	a.Logger.Info("openagent-control-plane starting", "port", a.Config.Port)
	return nil
}
