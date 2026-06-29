package repo

import (
	"context"
	"sync"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-sandbox/internal/domain"
)

// MemoryRepository is an in-memory implementation of Repository for development.
type MemoryRepository struct {
	mu       sync.RWMutex
	sandboxes map[string]*domain.Sandbox
}

// NewMemoryRepository creates a new in-memory repository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		sandboxes: make(map[string]*domain.Sandbox),
	}
}

// FindByID retrieves a sandbox by ID.
func (r *MemoryRepository) FindByID(ctx context.Context, id string) (*domain.Sandbox, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	sb, ok := r.sandboxes[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return sb, nil
}

// List returns all sandboxes.
func (r *MemoryRepository) List(ctx context.Context) ([]*domain.Sandbox, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*domain.Sandbox, 0, len(r.sandboxes))
	for _, sb := range r.sandboxes {
		result = append(result, sb)
	}
	return result, nil
}

// Save persists a sandbox.
func (r *MemoryRepository) Save(ctx context.Context, sb *domain.Sandbox) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.sandboxes[sb.ID] = sb
	return nil
}

// Delete removes a sandbox.
func (r *MemoryRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.sandboxes, id)
	return nil
}
