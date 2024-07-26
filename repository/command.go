package repository

import "context"

type Command[M any] interface {
	Exec(ctx context.Context, model M) error
}
