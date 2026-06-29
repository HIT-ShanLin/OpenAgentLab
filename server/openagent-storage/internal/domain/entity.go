// Package domain defines the core entities for the storage module.
//
// Storage manages ephemeral block devices and shared volumes
// dedicated to sandbox environments.
package domain

import "fmt"

// Volume represents a storage volume attached to a sandbox.
type Volume struct {
	ID        string
	Type      VolumeType
	SizeBytes int64
	SandboxID string
	Status    VolumeStatus
}

// VolumeType specifies the storage backend.
type VolumeType string

const (
	TypeBlock    VolumeType = "block"
	TypeOverlayFS VolumeType = "overlayfs"
	TypeShared   VolumeType = "shared"
)

// VolumeStatus represents lifecycle state.
type VolumeStatus string

const (
	VolumePending  VolumeStatus = "pending"
	VolumeAttached VolumeStatus = "attached"
	VolumeDetached VolumeStatus = "detached"
)

// NewVolume creates a new Volume with validation.
func NewVolume(id string, vt VolumeType, sizeBytes int64, sandboxID string) (*Volume, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidVolume)
	}
	if sandboxID == "" {
		return nil, fmt.Errorf("%w: sandboxID is required", ErrInvalidVolume)
	}
	return &Volume{
		ID:        id,
		Type:      vt,
		SizeBytes: sizeBytes,
		SandboxID: sandboxID,
		Status:    VolumePending,
	}, nil
}

// Attach marks the volume as attached.
func (v *Volume) Attach() error {
	if v.Status != VolumePending {
		return fmt.Errorf("%w: cannot attach volume in status %s", ErrInvalidTransition, v.Status)
	}
	v.Status = VolumeAttached
	return nil
}
