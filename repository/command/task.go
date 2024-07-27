package command

import (
	"context"

	"github.com/rinonkia/cqs-repository-pattarn/model"
	"github.com/rinonkia/cqs-repository-pattarn/repository/record"
	"gorm.io/gorm"
)

type (
	putTaskCommand    struct{ db *gorm.DB }
	deleteTaskCommand struct{ db *gorm.DB }
)

func NewPutTaskCommand(db *gorm.DB) *putTaskCommand {
	return &putTaskCommand{db: db}
}

func (c *putTaskCommand) Exec(ctx context.Context, task *model.Task) error {
	rec, err := record.FromTask(task)
	if err != nil {
		return err
	}

	tx := c.db.WithContext(ctx)
	if tx.Where("id = ?", rec.ID).Updates(rec).RowsAffected == 0 {
		if err := tx.Create(rec).Error; err != nil {
			return err
		}
	}
	return nil
}

func NewDeleteTaskCommand(db *gorm.DB) *deleteTaskCommand {
	return &deleteTaskCommand{db: db}
}

func (c *deleteTaskCommand) Exec(ctx context.Context, id int) error {
	if err := c.db.WithContext(ctx).Delete(&record.Task{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
