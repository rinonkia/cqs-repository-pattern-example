package model

import "fmt"

type (
	Task struct {
		ID       int
		Name     string
		Priority Priority
		Status   Status
	}

	Status   string
	Priority string
)

const (
	StatusNotStarted Status = "not_started"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
)

const (
	PriorityHigh   Priority = "high"
	PriorityMiddle Priority = "middle"
	PriorityLow    Priority = "low"
)

func (s Status) String() string { return string(s) }

func StatusFromString(s string) (Status, error) {
	switch s {
	case StatusNotStarted.String():
		return StatusNotStarted, nil
	case StatusInProgress.String():
		return StatusInProgress, nil
	case StatusCompleted.String():
		return StatusCompleted, nil
	default:
		return "", fmt.Errorf("unknown status. value: %s ", s)
	}
}

func (p Priority) String() string { return string(p) }

func PriorityFromString(s string) (Priority, error) {
	switch s {
	case PriorityHigh.String():
		return PriorityHigh, nil
	case PriorityMiddle.String():
		return PriorityMiddle, nil
	case PriorityLow.String():
		return PriorityLow, nil
	default:
		return "", fmt.Errorf("unknown priority. value: %s", s)
	}
}
