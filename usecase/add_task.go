package usecase

import (
	"context"

	"github.com/rinonkia/cqs-repository-pattarn/model"
	"github.com/rinonkia/cqs-repository-pattarn/repository"
)

type (
	AddTaskUsecase struct {
		putTask repository.Command[*model.Task]
	}

	AddTaskUsecaseDTO struct {
		Name     string
		Priority string
	}

	AddTaskUsecaseResult struct {
		Err error
	}
)

func NewAddTaskUsecase(putTask repository.Command[*model.Task]) *AddTaskUsecase {
	return &AddTaskUsecase{
		putTask: putTask,
	}
}

func (uc *AddTaskUsecase) Exec(ctx context.Context, dto *AddTaskUsecaseDTO) *AddTaskUsecaseResult {
	p, err := model.PriorityFromString(dto.Priority)
	if err != nil {
		return &AddTaskUsecaseResult{Err: err}
	}

	err = uc.putTask.Command(ctx, &model.Task{
		Name:     dto.Name,
		Priority: p,
		Status:   model.StatusNotStarted,
	})
	if err != nil {
		return &AddTaskUsecaseResult{Err: err}
	}

	return &AddTaskUsecaseResult{Err: nil}
}
