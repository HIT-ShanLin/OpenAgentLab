package domain

import "errors"

var (
	ErrNotFound          = errors.New("network not found")
	ErrInvalidNetwork    = errors.New("invalid network")
	ErrInvalidTransition = errors.New("invalid state transition")
)
