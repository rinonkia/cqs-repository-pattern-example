package usecase

import (
	"context"

	"github.com/rinonkia/cqs-repository-pattarn/model"
	"github.com/rinonkia/cqs-repository-pattarn/repository"
)

type (
	GetTaskUsecase struct {
		getTask repository.Query[int, *model.Task]
	}

	GetTaskUsecaseDTO struct {
		ID     int
		Status string
	}

	GetTaskUsecaseResult struct {
		Err  error
		Data *model.Task
	}
)

func NewGetTaskUsecase(getTask repository.Query[int, *model.Task]) *GetTaskUsecase {
	return &GetTaskUsecase{getTask: getTask}
}

func (uc *GetTaskUsecase) Exec(ctx context.Context, dto *GetTaskUsecaseDTO) *GetTaskUsecaseResult {
	task, err := uc.getTask.Query(ctx, dto.ID)
	if err != nil {
		return &GetTaskUsecaseResult{Err: err}
	}

	return &GetTaskUsecaseResult{Data: task}
}
