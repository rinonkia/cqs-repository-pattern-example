package usecase

import (
	"context"

	"github.com/rinonkia/cqs-repository-pattarn/model"
	"github.com/rinonkia/cqs-repository-pattarn/repository"
)

type (
	UpdateTaskUsecase struct {
		getTask repository.Query[int, *model.Task]
		putTask repository.Command[*model.Task]
	}

	UpdateTaskUsecaseDTO struct {
		ID       int
		Name     string
		Priority string
		Status   string
	}

	UpdateTaskUsecaseResult struct {
		Err error
	}
)

func NewUpdateTaskUsecase(
	getTask repository.Query[int, *model.Task],
	putTask repository.Command[*model.Task],
) *UpdateTaskUsecase {
	return &UpdateTaskUsecase{
		getTask: getTask,
		putTask: putTask,
	}
}

func (uc *UpdateTaskUsecase) Exec(ctx context.Context, dto *UpdateTaskUsecaseDTO) *UpdateTaskUsecaseResult {
	task, err := uc.getTask.Query(ctx, dto.ID)
	if err != nil {
		return &UpdateTaskUsecaseResult{Err: err}
	}

	status, err := model.StatusFromString(dto.Status)
	if err != nil {
		return &UpdateTaskUsecaseResult{Err: err}
	}

	priority, err := model.PriorityFromString(dto.Priority)
	if err != nil {
		return &UpdateTaskUsecaseResult{Err: err}
	}

	newTask := model.Task{
		ID:       task.ID,
		Name:     dto.Name,
		Priority: priority,
		Status:   status,
	}

	err = uc.putTask.Command(ctx, &newTask)
	if err != nil {
		return &UpdateTaskUsecaseResult{Err: err}
	}

	return &UpdateTaskUsecaseResult{Err: nil}
}
