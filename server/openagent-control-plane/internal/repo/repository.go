package repo

import (
	"context"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-control-plane/internal/domain"
)

// TaskRepository defines the persistence contract for task storage.
type TaskRepository interface {
	FindByID(ctx context.Context, id string) (*domain.Task, error)
	List(ctx context.Context) ([]*domain.Task, error)
	Save(ctx context.Context, t *domain.Task) error
	Delete(ctx context.Context, id string) error
}

// QueueRepository defines the persistence contract for task queues.
type QueueRepository interface {
	Enqueue(ctx context.Context, qName string, t *domain.Task) error
	Dequeue(ctx context.Context, qName string) (*domain.Task, error)
	Len(ctx context.Context, qName string) (int, error)
}
