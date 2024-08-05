package usecase

import (
	"context"

	"github.com/rinonkia/cqs-repository-pattarn/model"
	"github.com/rinonkia/cqs-repository-pattarn/repository"
)

type (
	GetTasksUsecase struct {
		getAllTasks      repository.QueryWithoutParam[[]*model.Task]
		getTasksByStatus repository.Query[model.Status, []*model.Task]
	}

	GetTasksUsecaseDTO struct {
		Status string
	}

	GetTasksUsecaseResult struct {
		Err  error
		Data []*model.Task
	}
)

func NewGetTasksUsecase(
	getAllTasks repository.QueryWithoutParam[[]*model.Task],
	getTasksByStatus repository.Query[model.Status, []*model.Task],
) *GetTasksUsecase {
	return &GetTasksUsecase{
		getAllTasks:      getAllTasks,
		getTasksByStatus: getTasksByStatus,
	}
}

func (uc *GetTasksUsecase) Exec(ctx context.Context, dto *GetTasksUsecaseDTO) *GetTasksUsecaseResult {
	switch dto.Status {
	case "":
		tasks, err := uc.getAllTasks.Query(ctx)
		if err != nil {
			return &GetTasksUsecaseResult{Err: err}
		}
		return &GetTasksUsecaseResult{Data: tasks}

	default:
		status, err := model.StatusFromString(dto.Status)
		if err != nil {
			return &GetTasksUsecaseResult{Err: err}
		}

		tasks, err := uc.getTasksByStatus.Query(ctx, status)
		if err != nil {
			return &GetTasksUsecaseResult{Err: err}
		}
		return &GetTasksUsecaseResult{Data: tasks}
	}
}
