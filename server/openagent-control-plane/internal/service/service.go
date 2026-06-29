package service

import (
	"context"
	"fmt"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-control-plane/internal/domain"
	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-control-plane/internal/repo"
)

// ControlPlaneService orchestrates task scheduling and routing.
type ControlPlaneService struct {
	repo      repo.TaskRepository
	queueRepo repo.QueueRepository
	scheduler domain.Scheduler
}

// New creates a new ControlPlaneService.
func New(taskRepo repo.TaskRepository, queueRepo repo.QueueRepository, s domain.Scheduler) *ControlPlaneService {
	return &ControlPlaneService{
		repo:      taskRepo,
		queueRepo: queueRepo,
		scheduler:  s,
	}
}

// Submit enqueues a new agent task.
func (s *ControlPlaneService) Submit(ctx context.Context, id, agentID, payload string) (*domain.Task, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	t, err := domain.NewTask(id, agentID, payload)
	if err != nil {
		return nil, fmt.Errorf("controlplane: submit: %w", err)
	}
	if err := s.repo.Save(ctx, t); err != nil {
		return nil, fmt.Errorf("controlplane: save: %w", err)
	}
	return t, nil
}

// Schedule dispatches a queued task to an available sandbox.
func (s *ControlPlaneService) Schedule(ctx context.Context, taskID string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	t, err := s.repo.FindByID(ctx, taskID)
	if err != nil {
		return fmt.Errorf("controlplane: find: %w", err)
	}
	sandboxID, err := s.scheduler.Assign(ctx, t)
	if err != nil {
		return fmt.Errorf("controlplane: schedule: %w", err)
	}
	if err := t.Schedule(sandboxID); err != nil {
		return err
	}
	return s.repo.Save(ctx, t)
}

// Get retrieves a task by ID.
func (s *ControlPlaneService) Get(ctx context.Context, id string) (*domain.Task, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, id)
}
