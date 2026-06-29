package repo

import (
	"context"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-network/internal/domain"
)

// Repository defines the persistence contract for network storage.
type Repository interface {
	FindByID(ctx context.Context, id string) (*domain.VirtualNetwork, error)
	List(ctx context.Context) ([]*domain.VirtualNetwork, error)
	Save(ctx context.Context, vn *domain.VirtualNetwork) error
	Delete(ctx context.Context, id string) error
}
