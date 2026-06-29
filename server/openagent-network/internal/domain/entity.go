// Package domain defines the core entities for the network module.
//
// Network manages virtual networking for sandbox environments,
// including veth pairs, bridges, and iptables rules.
package domain

import "fmt"

// VirtualNetwork represents an isolated virtual network for sandboxes.
type VirtualNetwork struct {
	ID      string
	Subnet  string
	Gateway string
	Status  NetworkStatus
}

// NetworkStatus represents lifecycle state.
type NetworkStatus string

const (
	NetworkPending NetworkStatus = "pending"
	NetworkActive  NetworkStatus = "active"
	NetworkDown    NetworkStatus = "down"
)

// Endpoint represents a network endpoint attached to a sandbox.
type Endpoint struct {
	ID        string
	NetworkID string
	SandboxID string
	IP        string
	MAC       string
}

// NewVirtualNetwork creates a new VirtualNetwork with validation.
func NewVirtualNetwork(id, subnet, gateway string) (*VirtualNetwork, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidNetwork)
	}
	return &VirtualNetwork{
		ID:      id,
		Subnet:  subnet,
		Gateway: gateway,
		Status:  NetworkPending,
	}, nil
}

// Activate marks the network as active.
func (vn *VirtualNetwork) Activate() error {
	if vn.Status != NetworkPending {
		return fmt.Errorf("%w: cannot activate network in status %s", ErrInvalidTransition, vn.Status)
	}
	vn.Status = NetworkActive
	return nil
}
