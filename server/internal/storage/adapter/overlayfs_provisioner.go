package adapter

import (
	"context"
	"fmt"
	"log"

	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/storage/domain"
)

// OverlayFSProvisioner implements domain.Provisioner using OverlayFS.
// Stub: will create overlay mounts on the host filesystem.
type OverlayFSProvisioner struct {
	BasePath string
}

// NewOverlayFSProvisioner creates a new OverlayFS provisioner.
func NewOverlayFSProvisioner(basePath string) *OverlayFSProvisioner {
	return &OverlayFSProvisioner{BasePath: basePath}
}

// Provision creates the overlay filesystem for a volume.
func (p *OverlayFSProvisioner) Provision(ctx context.Context, v *domain.Volume) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	log.Printf("[storage provisioner] provisioning volume %s (type=%s, size=%d) at %s",
		v.ID, v.Type, v.SizeBytes, p.BasePath)
	// TODO: create overlay mount via mount(2) syscall
	return fmt.Errorf("overlayfs provisioner: not yet implemented")
}
