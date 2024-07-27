package usecase

import (
	"context"

	"github.com/rinonkia/cqs-repository-pattarn/model"
	"github.com/rinonkia/cqs-repository-pattarn/repository"
)

type (
	DeleteTaskUsecase struct {
		getTasksQuery     repository.Query[int, *model.Task]
		deleteTaskCommand repository.Command[int]
	}

	DeleteTaskUsecaseDTO struct {
		ID int
	}

	DeleteTaskUsecaseResult struct {
		Err error
	}
)

func NewDeleteTaskUsecase(
	getTaskQuery repository.Query[int, *model.Task],
	deleteTaskCommand repository.Command[int],
) *DeleteTaskUsecase {
	return &DeleteTaskUsecase{
		getTasksQuery:     getTaskQuery,
		deleteTaskCommand: deleteTaskCommand,
	}
}

func (uc *DeleteTaskUsecase) Exec(ctx context.Context, dto *DeleteTaskUsecaseDTO) *DeleteTaskUsecaseResult {
	_, err := uc.getTasksQuery.Exec(ctx, dto.ID)
	if err != nil {
		return &DeleteTaskUsecaseResult{Err: err}
	}

	if err := uc.deleteTaskCommand.Exec(ctx, dto.ID); err != nil {
		return &DeleteTaskUsecaseResult{Err: err}
	}

	return &DeleteTaskUsecaseResult{Err: nil}
}
