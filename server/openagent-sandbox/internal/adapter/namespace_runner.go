package adapter

import (
	"context"
	"fmt"
	"log"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-sandbox/internal/domain"
)

// NamespaceRunner implements domain.Runner using Linux namespace primitives.
// This is a stub that logs intent; production implementation will invoke
// clone(2) with CLONE_NEW* flags or delegate to a container runtime.
type NamespaceRunner struct{}

// NewNamespaceRunner creates a new namespace-based runner.
func NewNamespaceRunner() *NamespaceRunner {
	return &NamespaceRunner{}
}

// Run starts the sandbox using Linux namespace isolation.
func (r *NamespaceRunner) Run(ctx context.Context, sb *domain.Sandbox) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	log.Printf("[sandbox runner] starting sandbox %s with namespaces %v, cgroups %+v",
		sb.ID, sb.Namespaces, sb.Cgroups)
	// TODO: implement clone(2) + exec or delegate to runc
	return fmt.Errorf("namespace runner: not yet implemented")
}
