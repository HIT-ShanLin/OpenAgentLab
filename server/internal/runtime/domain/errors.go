package domain

import "errors"

var (
	ErrNotFound          = errors.New("runtime not found")
	ErrInvalidRuntime    = errors.New("invalid runtime")
	ErrInvalidTransition = errors.New("invalid state transition")
)
