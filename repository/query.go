package repository

import "context"

type Query[P, M any] interface {
	Query(ctx context.Context, param P) (M, error)
}

// QueryWithoutParam for query all rows
type QueryWithoutParam[M any] interface {
	Query(ctx context.Context) (M, error)
}
