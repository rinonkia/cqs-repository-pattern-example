package usecase

import (
	"context"

	"github.com/rinonkia/cqs-repository-pattarn/repository"
)

type (
	DeleteTaskUsecase struct {
		deleteTaskCommand repository.Command[int]
	}

	DeleteTaskUsecaseDTO struct {
		ID int
	}

	DeleteTaskUsecaseResult struct {
		Err error
	}
)

func NewDeleteTaskUsecase(deleteTaskCommand repository.Command[int]) *DeleteTaskUsecase {
	return &DeleteTaskUsecase{deleteTaskCommand: deleteTaskCommand}
}

func (uc *DeleteTaskUsecase) Exec(ctx context.Context, dto *DeleteTaskUsecaseDTO) *DeleteTaskUsecaseResult {
	err := uc.deleteTaskCommand.Exec(ctx, dto.ID)
	if err != nil {
		return &DeleteTaskUsecaseResult{Err: err}
	}

	return &DeleteTaskUsecaseResult{Err: nil}
}
