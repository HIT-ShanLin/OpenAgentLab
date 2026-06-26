package adapter

import (
	"context"
	"fmt"
	"log"

	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/network/domain"
)

// LinuxBridge implements domain.Driver using Linux bridge + veth + iptables.
// Stub: will invoke ip link, iptables, and netlink syscalls.
type LinuxBridge struct {
	BridgeName string
}

// NewLinuxBridge creates a new linux bridge driver.
func NewLinuxBridge(bridgeName string) *LinuxBridge {
	return &LinuxBridge{BridgeName: bridgeName}
}

// Setup creates the bridge and configures the virtual network.
func (d *LinuxBridge) Setup(ctx context.Context, vn *domain.VirtualNetwork) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	log.Printf("[network driver] setting up virtual network %s (subnet=%s, gateway=%s, bridge=%s)",
		vn.ID, vn.Subnet, vn.Gateway, d.BridgeName)
	// TODO: create bridge via netlink, set up iptables NAT rules
	return fmt.Errorf("linux bridge driver: not yet implemented")
}
