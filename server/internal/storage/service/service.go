package service

import (
	"context"
	"fmt"

	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/storage/domain"
	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/storage/repo"
)

// VolumeService manages the lifecycle of sandbox storage volumes.
type VolumeService struct {
	repo      repo.Repository
	provision domain.Provisioner
}

// New creates a new VolumeService.
func New(r repo.Repository, p domain.Provisioner) *VolumeService {
	return &VolumeService{repo: r, provision: p}
}

// Create provisions a new volume.
func (s *VolumeService) Create(ctx context.Context, id string, vt domain.VolumeType, sizeBytes int64, sandboxID string) (*domain.Volume, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	v, err := domain.NewVolume(id, vt, sizeBytes, sandboxID)
	if err != nil {
		return nil, fmt.Errorf("volume service: create: %w", err)
	}
	if err := s.provision.Provision(ctx, v); err != nil {
		return nil, fmt.Errorf("volume service: provision: %w", err)
	}
	if err := s.repo.Save(ctx, v); err != nil {
		return nil, fmt.Errorf("volume service: save: %w", err)
	}
	return v, nil
}

// Attach connects a volume to its sandbox.
func (s *VolumeService) Attach(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	v, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("volume service: find: %w", err)
	}
	if err := v.Attach(); err != nil {
		return err
	}
	return s.repo.Save(ctx, v)
}

// Get retrieves a volume by ID.
func (s *VolumeService) Get(ctx context.Context, id string) (*domain.Volume, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, id)
}
