package repo

import (
	"context"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-runtime/internal/domain"
)

// Repository defines the persistence contract for runtime storage.
type Repository interface {
	FindByID(ctx context.Context, id string) (*domain.Runtime, error)
	List(ctx context.Context) ([]*domain.Runtime, error)
	Save(ctx context.Context, r *domain.Runtime) error
	Delete(ctx context.Context, id string) error
}
