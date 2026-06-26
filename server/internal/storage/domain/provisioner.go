package domain

import "context"

// Provisioner abstracts the mechanism to create storage volumes.
type Provisioner interface {
	Provision(ctx context.Context, v *Volume) error
}
