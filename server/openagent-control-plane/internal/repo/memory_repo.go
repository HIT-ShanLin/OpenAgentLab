package repo

import (
	"context"
	"sync"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-control-plane/internal/domain"
)

// MemoryTaskRepository is an in-memory implementation of TaskRepository.
type MemoryTaskRepository struct {
	mu    sync.RWMutex
	tasks map[string]*domain.Task
}

// NewMemoryTaskRepository creates a new in-memory task repository.
func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{tasks: make(map[string]*domain.Task)}
}

func (r *MemoryTaskRepository) FindByID(ctx context.Context, id string) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.tasks[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return t, nil
}

func (r *MemoryTaskRepository) List(ctx context.Context) ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*domain.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		result = append(result, t)
	}
	return result, nil
}

func (r *MemoryTaskRepository) Save(ctx context.Context, t *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[t.ID] = t
	return nil
}

func (r *MemoryTaskRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.tasks, id)
	return nil
}

// MemoryQueueRepository is an in-memory implementation of QueueRepository.
type MemoryQueueRepository struct {
	mu     sync.RWMutex
	queues map[string][]*domain.Task
}

// NewMemoryQueueRepository creates a new in-memory queue repository.
func NewMemoryQueueRepository() *MemoryQueueRepository {
	return &MemoryQueueRepository{queues: make(map[string][]*domain.Task)}
}

func (r *MemoryQueueRepository) Enqueue(ctx context.Context, qName string, t *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.queues[qName] = append(r.queues[qName], t)
	return nil
}

func (r *MemoryQueueRepository) Dequeue(ctx context.Context, qName string) (*domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	q := r.queues[qName]
	if len(q) == 0 {
		return nil, domain.ErrNotFound
	}
	t := q[0]
	r.queues[qName] = q[1:]
	return t, nil
}

func (r *MemoryQueueRepository) Len(ctx context.Context, qName string) (int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.queues[qName]), nil
}
