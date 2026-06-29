// Package domain defines the core entities and rules for the sandbox module.
//
// A Sandbox represents an isolated execution environment for running untrusted
// Agent code, built on Linux Namespace and cgroup primitives.
package domain

import "fmt"

// Sandbox is the central entity representing an isolated execution environment.
type Sandbox struct {
	ID         string
	Status     SandboxStatus
	Namespaces []Namespace
	Cgroups    CgroupLimits
	Image      string
}

// SandboxStatus represents the lifecycle state of a sandbox.
type SandboxStatus string

const (
	StatusPending   SandboxStatus = "pending"
	StatusRunning   SandboxStatus = "running"
	StatusStopped   SandboxStatus = "stopped"
	StatusDestroyed SandboxStatus = "destroyed"
)

// Namespace represents a Linux namespace type to isolate.
type Namespace string

const (
	NamespacePID   Namespace = "pid"
	NamespaceNet   Namespace = "net"
	NamespaceMnt   Namespace = "mnt"
	NamespaceUTS   Namespace = "uts"
	NamespaceIPC   Namespace = "ipc"
	NamespaceUser  Namespace = "user"
	NamespaceCgroup Namespace = "cgroup"
)

// CgroupLimits defines resource constraints for a sandbox.
type CgroupLimits struct {
	MemoryBytes int64
	CPUQuota    int64
	PidsLimit   int64
}

// NewSandbox creates a new Sandbox with required fields validated.
func NewSandbox(id, image string) (*Sandbox, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidSandbox)
	}
	if image == "" {
		return nil, fmt.Errorf("%w: image is required", ErrInvalidSandbox)
	}
	return &Sandbox{
		ID:     id,
		Status: StatusPending,
		Image:  image,
	}, nil
}

// AddNamespace adds a namespace isolation type to the sandbox.
func (s *Sandbox) AddNamespace(ns Namespace) {
	s.Namespaces = append(s.Namespaces, ns)
}

// SetCgroupLimits applies resource limits.
func (s *Sandbox) SetCgroupLimits(limits CgroupLimits) {
	s.Cgroups = limits
}

// Start transitions the sandbox to running state.
func (s *Sandbox) Start() error {
	if s.Status != StatusPending {
		return fmt.Errorf("%w: cannot start sandbox in status %s", ErrInvalidTransition, s.Status)
	}
	s.Status = StatusRunning
	return nil
}

// Stop transitions the sandbox to stopped state.
func (s *Sandbox) Stop() error {
	if s.Status != StatusRunning {
		return fmt.Errorf("%w: cannot stop sandbox in status %s", ErrInvalidTransition, s.Status)
	}
	s.Status = StatusStopped
	return nil
}
