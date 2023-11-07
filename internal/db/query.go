package db

import (
	"context"
	"github.com/scylladb/gocqlx/v2"
)

type QueryXService interface {
	BindStruct(arg interface{}) QueryXService
	BindMap(arg map[string]interface{}) QueryXService
	Exec(ctx context.Context) error
	SelectRelease(ctx context.Context, dest interface{}) error
	GetRelease(ctx context.Context, dest interface{}) error
}

type QueryXServiceImpl struct{
	QueryX *gocqlx.Queryx
}

func (q *QueryXServiceImpl) BindStruct(arg interface{}) QueryXService{
	return &QueryXServiceImpl{
		QueryX: q.QueryX.BindStruct(arg),
	}
}

func (q *QueryXServiceImpl) BindMap(arg map[string]interface{}) QueryXService{
	return &QueryXServiceImpl{
		QueryX: q.QueryX.BindMap(arg),
	}
}

func (q *QueryXServiceImpl) Exec(ctx context.Context) error{
	return  q.QueryX.Exec()
}

func (q *QueryXServiceImpl) SelectRelease(ctx context.Context, dest interface{}) error{
	return q.QueryX.SelectRelease(dest)
}

func (q *QueryXServiceImpl) GetRelease(ctx context.Context, dest interface{}) error{
	return q.QueryX.GetRelease(dest)
}