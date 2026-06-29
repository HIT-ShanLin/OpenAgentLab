// Package domain defines the core entities for the runtime module.
//
// Runtime manages the lifecycle of agent execution environments,
// supporting both container-based and MicroVM-based isolation.
package domain

import "fmt"

// Runtime is the execution environment that runs agent code.
type Runtime struct {
	ID       string
	Type     RuntimeType
	Status   RuntimeStatus
	SandboxID string
}

// RuntimeType specifies the isolation technology.
type RuntimeType string

const (
	TypeContainer RuntimeType = "container"
	TypeMicroVM   RuntimeType = "microvm"
)

// RuntimeStatus represents lifecycle state.
type RuntimeStatus string

const (
	RuntimePending  RuntimeStatus = "pending"
	RuntimeRunning  RuntimeStatus = "running"
	RuntimeExited   RuntimeStatus = "exited"
	RuntimeFailed   RuntimeStatus = "failed"
)

// NewRuntime creates a new Runtime with validation.
func NewRuntime(id string, rt RuntimeType, sandboxID string) (*Runtime, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidRuntime)
	}
	if sandboxID == "" {
		return nil, fmt.Errorf("%w: sandboxID is required", ErrInvalidRuntime)
	}
	return &Runtime{
		ID:        id,
		Type:      rt,
		Status:    RuntimePending,
		SandboxID: sandboxID,
	}, nil
}

// Start marks the runtime as running.
func (r *Runtime) Start() error {
	if r.Status != RuntimePending {
		return fmt.Errorf("%w: cannot start runtime in status %s", ErrInvalidTransition, r.Status)
	}
	r.Status = RuntimeRunning
	return nil
}
