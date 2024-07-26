package model

import "fmt"

type (
	Task struct {
		ID     int
		Name   string
		Status Status
	}

	Status string
)

const (
	NotStarted Status = "not_started"
	InProgress Status = "in_progress"
	Completed  Status = "completed"
)

func (s Status) String() string { return string(s) }

func StatusFromString(s string) (Status, error) {
	switch s {
	case NotStarted.String():
		return NotStarted, nil
	case InProgress.String():
		return InProgress, nil
	case Completed.String():
		return Completed, nil
	default:
		return "", fmt.Errorf("unknown status. value: %s ", s)
	}
}
