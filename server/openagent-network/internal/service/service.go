package service

import (
	"context"
	"fmt"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-network/internal/domain"
	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-network/internal/repo"
)

// NetworkService manages virtual network lifecycle.
type NetworkService struct {
	repo    repo.Repository
	driver  domain.Driver
}

// New creates a new NetworkService.
func New(r repo.Repository, d domain.Driver) *NetworkService {
	return &NetworkService{repo: r, driver: d}
}

// Create provisions a new virtual network.
func (s *NetworkService) Create(ctx context.Context, id, subnet, gateway string) (*domain.VirtualNetwork, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	net, err := domain.NewVirtualNetwork(id, subnet, gateway)
	if err != nil {
		return nil, fmt.Errorf("network service: create: %w", err)
	}
	if err := s.driver.Setup(ctx, net); err != nil {
		return nil, fmt.Errorf("network service: setup: %w", err)
	}
	net.Activate()
	if err := s.repo.Save(ctx, net); err != nil {
		return nil, fmt.Errorf("network service: save: %w", err)
	}
	return net, nil
}

// Get retrieves a virtual network by ID.
func (s *NetworkService) Get(ctx context.Context, id string) (*domain.VirtualNetwork, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, id)
}

// ConnectEndpoint attaches a sandbox endpoint to the network.
func (s *NetworkService) ConnectEndpoint(ctx context.Context, networkID, sandboxID, ip string) (*domain.Endpoint, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	ep := &domain.Endpoint{
		ID:        fmt.Sprintf("ep-%s-%s", sandboxID, networkID[:8]),
		NetworkID: networkID,
		SandboxID: sandboxID,
		IP:        ip,
	}
	return ep, nil
}
