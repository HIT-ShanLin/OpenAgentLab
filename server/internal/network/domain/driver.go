package domain

import "context"

// Driver abstracts the mechanism to set up virtual networks.
type Driver interface {
	Setup(ctx context.Context, vn *VirtualNetwork) error
}
