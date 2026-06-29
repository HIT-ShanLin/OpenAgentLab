package repo

import (
	"context"
	"sync"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-storage/internal/domain"
)

// MemoryRepository is an in-memory implementation of Repository for development.
type MemoryRepository struct {
	mu      sync.RWMutex
	volumes map[string]*domain.Volume
}

// NewMemoryRepository creates a new in-memory repository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{volumes: make(map[string]*domain.Volume)}
}

func (r *MemoryRepository) FindByID(ctx context.Context, id string) (*domain.Volume, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.volumes[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return v, nil
}

func (r *MemoryRepository) List(ctx context.Context) ([]*domain.Volume, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*domain.Volume, 0, len(r.volumes))
	for _, v := range r.volumes {
		result = append(result, v)
	}
	return result, nil
}

func (r *MemoryRepository) Save(ctx context.Context, v *domain.Volume) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.volumes[v.ID] = v
	return nil
}

func (r *MemoryRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.volumes, id)
	return nil
}
