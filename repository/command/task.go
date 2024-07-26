package command

import (
	"context"

	"github.com/rinonkia/sqs-repository-pattarn/model"
)

type (
	putTaskCommand    struct{}
	deleteTaskCommand struct{}
)

func NewPutTaskCommand() *putTaskCommand {
	return &putTaskCommand{}
}

func (c *putTaskCommand) Exec(ctx context.Context, task *model.Task) error {
	return nil
}

func NewDeleteTaskCommand() *deleteTaskCommand {
	return &deleteTaskCommand{}
}

func (c *deleteTaskCommand) Exec(ctx context.Context, id int) error {
	return nil
}
