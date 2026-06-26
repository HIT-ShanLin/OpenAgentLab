package adapter

import (
	"context"
	"fmt"
	"log"

	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/controlplane/domain"
)

// SimpleScheduler implements domain.Scheduler with a basic round-robin strategy.
// Stub: will be replaced with priority-based scheduling across sandbox pools.
type SimpleScheduler struct{}

// NewSimpleScheduler creates a new simple scheduler.
func NewSimpleScheduler() *SimpleScheduler {
	return &SimpleScheduler{}
}

// Assign selects an available sandbox for the task.
func (s *SimpleScheduler) Assign(ctx context.Context, t *domain.Task) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	log.Printf("[scheduler] assigning task %s (agent=%s, priority=%d)", t.ID, t.AgentID, t.Priority)
	// TODO: implement sandbox pool selection with resource awareness
	return "", fmt.Errorf("simple scheduler: no sandbox available (not yet implemented)")
}
