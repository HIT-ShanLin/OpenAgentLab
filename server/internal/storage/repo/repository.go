package repo

import (
	"context"

	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/storage/domain"
)

// Repository defines the persistence contract for volume storage.
type Repository interface {
	FindByID(ctx context.Context, id string) (*domain.Volume, error)
	List(ctx context.Context) ([]*domain.Volume, error)
	Save(ctx context.Context, v *domain.Volume) error
	Delete(ctx context.Context, id string) error
}
