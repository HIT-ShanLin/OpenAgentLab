// Package domain defines the core entities for the control plane module.
//
// ControlPlane orchestrates the scheduling, queuing, and routing
// of agent tasks across sandbox environments.
package domain

import "fmt"

// Task represents a unit of work to be executed in a sandbox.
type Task struct {
	ID        string
	AgentID   string
	SandboxID string
	Status    TaskStatus
	Priority  int
	Payload   string
}

// TaskStatus represents the lifecycle of a task.
type TaskStatus string

const (
	TaskQueued     TaskStatus = "queued"
	TaskScheduled  TaskStatus = "scheduled"
	TaskRunning    TaskStatus = "running"
	TaskCompleted  TaskStatus = "completed"
	TaskFailed     TaskStatus = "failed"
)

// Queue is a FIFO task queue for pending agent requests.
type Queue struct {
	Name     string
	Capacity int
}

// NewTask creates a new Task with validation.
func NewTask(id, agentID, payload string) (*Task, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: id is required", ErrInvalidTask)
	}
	if agentID == "" {
		return nil, fmt.Errorf("%w: agentID is required", ErrInvalidTask)
	}
	return &Task{
		ID:       id,
		AgentID:  agentID,
		Status:   TaskQueued,
		Priority: 0,
		Payload:  payload,
	}, nil
}

// Schedule assigns the task to a sandbox.
func (t *Task) Schedule(sandboxID string) error {
	if t.Status != TaskQueued {
		return fmt.Errorf("%w: cannot schedule task in status %s", ErrInvalidTransition, t.Status)
	}
	t.SandboxID = sandboxID
	t.Status = TaskScheduled
	return nil
}

// Complete marks the task as completed.
func (t *Task) Complete() error {
	if t.Status != TaskRunning {
		return fmt.Errorf("%w: cannot complete task in status %s", ErrInvalidTransition, t.Status)
	}
	t.Status = TaskCompleted
	return nil
}
