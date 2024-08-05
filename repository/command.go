package repository

import "context"

type Command[M any] interface {
	Command(ctx context.Context, model M) error
}
