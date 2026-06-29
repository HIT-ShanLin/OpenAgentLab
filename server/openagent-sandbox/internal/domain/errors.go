package domain

import "errors"

var (
	// ErrNotFound is returned when a sandbox cannot be found.
	ErrNotFound = errors.New("sandbox not found")

	// ErrInvalidSandbox is returned when sandbox fields are invalid.
	ErrInvalidSandbox = errors.New("invalid sandbox")

	// ErrInvalidTransition is returned when a state transition is not allowed.
	ErrInvalidTransition = errors.New("invalid state transition")

	// ErrSandboxExists is returned when creating a sandbox with a duplicate ID.
	ErrSandboxExists = errors.New("sandbox already exists")
)
