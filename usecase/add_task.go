package usecase

import (
	"context"

	"github.com/rinonkia/cqs-repository-pattarn/model"
	"github.com/rinonkia/cqs-repository-pattarn/repository"
)

type (
	AddTaskUsecase struct {
		putTaskCommand repository.Command[*model.Task]
	}

	AddTaskUsecaseDTO struct {
		Name   string
		Status string
	}

	AddTaskUsecaseResult struct {
		Err error
	}
)

func NewAddTaskUsecase(putTaskCommand repository.Command[*model.Task]) *AddTaskUsecase {
	return &AddTaskUsecase{
		putTaskCommand: putTaskCommand,
	}
}

func (uc *AddTaskUsecase) Execute(ctx context.Context, dto *AddTaskUsecaseDTO) *AddTaskUsecaseResult {
	err := uc.putTaskCommand.Exec(ctx, &model.Task{
		Name:   dto.Name,
		Status: model.StatusNotStarted,
	})
	if err != nil {
		return &AddTaskUsecaseResult{Err: err}
	}

	return &AddTaskUsecaseResult{Err: nil}
}
