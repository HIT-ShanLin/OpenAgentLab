package service

import (
	"context"
	"fmt"

	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/runtime/domain"
	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/runtime/repo"
)

// RuntimeService orchestrates agent execution environment lifecycle.
type RuntimeService struct {
	repo    repo.Repository
	starter domain.Starter
}

// New creates a new RuntimeService.
func New(r repo.Repository, s domain.Starter) *RuntimeService {
	return &RuntimeService{repo: r, starter: s}
}

// Create provisions a new execution environment.
func (s *RuntimeService) Create(ctx context.Context, id string, rt domain.RuntimeType, sandboxID string) (*domain.Runtime, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	rtd, err := domain.NewRuntime(id, rt, sandboxID)
	if err != nil {
		return nil, fmt.Errorf("runtime service: create: %w", err)
	}
	if err := s.repo.Save(ctx, rtd); err != nil {
		return nil, fmt.Errorf("runtime service: save: %w", err)
	}
	return rtd, nil
}

// Start launches an execution environment.
func (s *RuntimeService) Start(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	r, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("runtime service: find: %w", err)
	}
	if err := r.Start(); err != nil {
		return err
	}
	if err := s.repo.Save(ctx, r); err != nil {
		return fmt.Errorf("runtime service: save: %w", err)
	}
	return s.starter.Start(ctx, r)
}

// Get retrieves a runtime by ID.
func (s *RuntimeService) Get(ctx context.Context, id string) (*domain.Runtime, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, id)
}
