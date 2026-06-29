package repo

import (
	"context"
	"sync"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-runtime/internal/domain"
)

// MemoryRepository is an in-memory implementation of Repository for development.
type MemoryRepository struct {
	mu       sync.RWMutex
	runtimes map[string]*domain.Runtime
}

// NewMemoryRepository creates a new in-memory repository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		runtimes: make(map[string]*domain.Runtime),
	}
}

func (r *MemoryRepository) FindByID(ctx context.Context, id string) (*domain.Runtime, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	rt, ok := r.runtimes[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return rt, nil
}

func (r *MemoryRepository) List(ctx context.Context) ([]*domain.Runtime, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*domain.Runtime, 0, len(r.runtimes))
	for _, rt := range r.runtimes {
		result = append(result, rt)
	}
	return result, nil
}

func (r *MemoryRepository) Save(ctx context.Context, rt *domain.Runtime) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.runtimes[rt.ID] = rt
	return nil
}

func (r *MemoryRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.runtimes, id)
	return nil
}
