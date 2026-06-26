package service

import (
	"context"
	"fmt"

	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/sandbox/domain"
	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/sandbox/repo"
)

// SandboxService orchestrates sandbox lifecycle operations.
type SandboxService struct {
	repo   repo.Repository
	runner domain.Runner
}

// New creates a new SandboxService.
func New(r repo.Repository, runner domain.Runner) *SandboxService {
	return &SandboxService{repo: r, runner: runner}
}

// Create provisions a new sandbox with the given configuration.
func (s *SandboxService) Create(ctx context.Context, id, image string) (*domain.Sandbox, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	sb, err := domain.NewSandbox(id, image)
	if err != nil {
		return nil, fmt.Errorf("sandbox service: create: %w", err)
	}
	if err := s.repo.Save(ctx, sb); err != nil {
		return nil, fmt.Errorf("sandbox service: save: %w", err)
	}
	return sb, nil
}

// Start launches a sandbox by ID.
func (s *SandboxService) Start(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	sb, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("sandbox service: find: %w", err)
	}
	if err := sb.Start(); err != nil {
		return err
	}
	if err := s.repo.Save(ctx, sb); err != nil {
		return fmt.Errorf("sandbox service: save: %w", err)
	}
	return s.runner.Run(ctx, sb)
}

// Stop halts a running sandbox.
func (s *SandboxService) Stop(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	sb, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("sandbox service: find: %w", err)
	}
	if err := sb.Stop(); err != nil {
		return err
	}
	return s.repo.Save(ctx, sb)
}

// Get retrieves a sandbox by ID.
func (s *SandboxService) Get(ctx context.Context, id string) (*domain.Sandbox, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, id)
}

// List returns all sandboxes.
func (s *SandboxService) List(ctx context.Context) ([]*domain.Sandbox, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return s.repo.List(ctx)
}
