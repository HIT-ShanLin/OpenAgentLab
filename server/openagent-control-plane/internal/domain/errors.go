package domain

import "errors"

var (
	ErrNotFound          = errors.New("task not found")
	ErrInvalidTask       = errors.New("invalid task")
	ErrInvalidTransition = errors.New("invalid state transition")
	ErrQueueFull         = errors.New("queue is full")
)
