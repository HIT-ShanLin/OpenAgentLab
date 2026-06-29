package domain

import "context"

// Runner abstracts the underlying mechanism to execute a sandbox.
// Implementations may use Linux namespaces directly, runc, or Firecracker MicroVM.
type Runner interface {
	Run(ctx context.Context, sb *Sandbox) error
}
