package adapter

import (
	"context"
	"fmt"
	"log"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-runtime/internal/domain"
)

// ContainerStarter implements domain.Starter using a container runtime.
// Stub: will delegate to runc / containerd / Firecracker.
type ContainerStarter struct{}

// NewContainerStarter creates a new container-based starter.
func NewContainerStarter() *ContainerStarter {
	return &ContainerStarter{}
}

// Start launches the runtime environment.
func (s *ContainerStarter) Start(ctx context.Context, r *domain.Runtime) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	log.Printf("[runtime starter] starting runtime %s (type=%s, sandbox=%s)",
		r.ID, r.Type, r.SandboxID)
	// TODO: invoke container runtime or KVM
	return fmt.Errorf("container starter: not yet implemented")
}
