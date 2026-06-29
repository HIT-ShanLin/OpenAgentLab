package domain

import "errors"

var (
	ErrNotFound          = errors.New("volume not found")
	ErrInvalidVolume     = errors.New("invalid volume")
	ErrInvalidTransition = errors.New("invalid state transition")
)
