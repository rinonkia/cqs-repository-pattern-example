package query

import (
	"context"

	"github.com/rinonkia/cqs-repository-pattarn/model"
	"github.com/rinonkia/cqs-repository-pattarn/repository/record"
	"gorm.io/gorm"
)

type (
	getTaskByIDQuery      struct{ db *gorm.DB }
	getTasksByStatusQuery struct{ db *gorm.DB }

	getAllTasksQuery struct{ db *gorm.DB }
)

func NewGetTaskByIDQuery(db *gorm.DB) *getTaskByIDQuery {
	return &getTaskByIDQuery{db: db}
}

func (q *getTaskByIDQuery) Exec(ctx context.Context, id int) (*model.Task, error) {
	var rec *record.Task
	if err := q.db.WithContext(ctx).First(&rec, "id = ?", id).Error; err != nil {
		return nil, err
	}

	task, err := record.TaskFromRecord(rec)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func NewGetTasksByStatusQuery(db *gorm.DB) *getTasksByStatusQuery {
	return &getTasksByStatusQuery{db: db}
}

func (q *getTasksByStatusQuery) Exec(ctx context.Context, status model.Status) ([]*model.Task, error) {
	var recs []*record.Task
	tx := q.db.WithContext(ctx)
	if err := tx.Find(&recs, "status = ?", status.String()).Error; err != nil {
		return nil, err
	}

	tasks, err := record.TasksFromRecords(recs)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func NewGetAllTasksQuery(db *gorm.DB) *getAllTasksQuery {
	return &getAllTasksQuery{db: db}
}

func (q *getAllTasksQuery) Exec(ctx context.Context) ([]*model.Task, error) {
	var recs []*record.Task
	if err := q.db.WithContext(ctx).Find(&recs).Error; err != nil {
		return nil, err
	}

	tasks, err := record.TasksFromRecords(recs)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
