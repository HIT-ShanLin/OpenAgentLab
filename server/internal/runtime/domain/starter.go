package domain

import "context"

// Starter abstracts the mechanism to launch an execution environment.
type Starter interface {
	Start(ctx context.Context, r *Runtime) error
}
