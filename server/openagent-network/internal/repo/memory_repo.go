package repo

import (
	"context"
	"sync"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-network/internal/domain"
)

// MemoryRepository is an in-memory implementation of Repository for development.
type MemoryRepository struct {
	mu       sync.RWMutex
	networks map[string]*domain.VirtualNetwork
}

// NewMemoryRepository creates a new in-memory repository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{networks: make(map[string]*domain.VirtualNetwork)}
}

func (r *MemoryRepository) FindByID(ctx context.Context, id string) (*domain.VirtualNetwork, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	net, ok := r.networks[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return net, nil
}

func (r *MemoryRepository) List(ctx context.Context) ([]*domain.VirtualNetwork, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*domain.VirtualNetwork, 0, len(r.networks))
	for _, net := range r.networks {
		result = append(result, net)
	}
	return result, nil
}

func (r *MemoryRepository) Save(ctx context.Context, vn *domain.VirtualNetwork) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.networks[vn.ID] = vn
	return nil
}

func (r *MemoryRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.networks, id)
	return nil
}
