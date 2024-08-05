package usecase

import (
	"context"

	"github.com/rinonkia/cqs-repository-pattarn/model"
	"github.com/rinonkia/cqs-repository-pattarn/repository"
)

type (
	DeleteTaskUsecase struct {
		getTasks   repository.Query[int, *model.Task]
		deleteTask repository.Command[int]
	}

	DeleteTaskUsecaseDTO struct {
		ID int
	}

	DeleteTaskUsecaseResult struct {
		Err error
	}
)

func NewDeleteTaskUsecase(
	getTask repository.Query[int, *model.Task],
	deleteTask repository.Command[int],
) *DeleteTaskUsecase {
	return &DeleteTaskUsecase{
		getTasks:   getTask,
		deleteTask: deleteTask,
	}
}

func (uc *DeleteTaskUsecase) Exec(ctx context.Context, dto *DeleteTaskUsecaseDTO) *DeleteTaskUsecaseResult {
	_, err := uc.getTasks.Query(ctx, dto.ID)
	if err != nil {
		return &DeleteTaskUsecaseResult{Err: err}
	}

	if err := uc.deleteTask.Command(ctx, dto.ID); err != nil {
		return &DeleteTaskUsecaseResult{Err: err}
	}

	return &DeleteTaskUsecaseResult{Err: nil}
}
