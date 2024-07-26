package repository

import "context"

type Query[P, M any] interface {
	Exec(ctx context.Context, param P) (M, error)
}

// QueryWithoutParam for query all
type QueryWithoutParam[M any] interface {
	Exec(ctx context.Context) (M, error)
}
