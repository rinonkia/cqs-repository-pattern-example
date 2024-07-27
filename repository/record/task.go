package record

import (
	"github.com/rinonkia/cqs-repository-pattarn/model"
)

type Task struct {
	ID       uint
	Name     string
	Priority string
	Status   string
}

func FromTask(task *model.Task) (*Task, error) {
	return &Task{
		ID:       uint(task.ID),
		Name:     task.Name,
		Priority: task.Priority.String(),
		Status:   task.Status.String(),
	}, nil
}

func TaskFromRecord(rec *Task) (*model.Task, error) {
	priority, err := model.PriorityFromString(rec.Priority)
	if err != nil {
		return nil, err
	}

	status, err := model.StatusFromString(rec.Status)
	if err != nil {
		return nil, err
	}

	return &model.Task{
		ID:       int(rec.ID),
		Name:     rec.Name,
		Priority: priority,
		Status:   status,
	}, nil
}

func TasksFromRecords(recs []*Task) ([]*model.Task, error) {
	tasks := make([]*model.Task, len(recs))

	for i, rec := range recs {
		task, err := TaskFromRecord(rec)
		if err != nil {
			return nil, err
		}
		tasks[i] = task
	}
	return tasks, nil
}
