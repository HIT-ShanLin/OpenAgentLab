package domain

import "context"

// Scheduler abstracts the task-to-sandbox assignment logic.
type Scheduler interface {
	Assign(ctx context.Context, t *Task) (sandboxID string, err error)
}
