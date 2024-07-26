package query

import (
	"context"

	"github.com/rinonkia/cqs-repository-pattarn/model"
)

type (
	getTaskByIDQuery      struct{}
	getTasksByStatusQuery struct{}

	getAllTasksQuery struct{}
)

func NewGetTaskByIDQuery() *getTaskByIDQuery {
	return &getTaskByIDQuery{}
}

func (q *getTaskByIDQuery) Exec(ctx context.Context, id int) (*model.Task, error) {
	return nil, nil
}

func NewGetTasksByStatusQuery() *getTasksByStatusQuery {
	return &getTasksByStatusQuery{}
}

func (q *getTasksByStatusQuery) Exec(ctx context.Context, status model.Status) ([]*model.Task, error) {
	return nil, nil
}

func NewGetAllTasksQuery() *getAllTasksQuery {
	return &getAllTasksQuery{}
}

func (q *getAllTasksQuery) Exec(ctx context.Context) ([]*model.Task, error) {
	return nil, nil
}
