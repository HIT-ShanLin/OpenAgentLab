package repo

import (
	"context"

	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/sandbox/domain"
)

// Repository defines the persistence contract for sandbox storage.
type Repository interface {
	FindByID(ctx context.Context, id string) (*domain.Sandbox, error)
	List(ctx context.Context) ([]*domain.Sandbox, error)
	Save(ctx context.Context, sb *domain.Sandbox) error
	Delete(ctx context.Context, id string) error
}
