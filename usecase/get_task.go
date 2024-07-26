package usecase

import (
	"context"

	"github.com/rinonkia/cqs-repository-pattarn/model"
	"github.com/rinonkia/cqs-repository-pattarn/repository"
)

type (
	GetTaskUsecase struct {
		getTaskQuery repository.Query[int, *model.Task]
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

func NewGetTaskUsecase(getTaskQuery repository.Query[int, *model.Task]) *GetTaskUsecase {
	return &GetTaskUsecase{getTaskQuery: getTaskQuery}
}

func (uc *GetTaskUsecase) Exec(ctx context.Context, dto *GetTaskUsecaseDTO) *GetTaskUsecaseResult {
	task, err := uc.getTaskQuery.Exec(ctx, dto.ID)
	if err != nil {
		return &GetTaskUsecaseResult{Err: err}
	}

	return &GetTaskUsecaseResult{Data: task}
}
