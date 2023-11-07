package dao

import (
	"context"

	daomodels "github.com/go-grpc-service/internal/dao/dao_models"
	"github.com/go-grpc-service/internal/db"
	"github.com/scylladb/gocqlx/v2/table"
)


var(
	movieColumns = []string{"movie_id", "name", "genre", "description", "rating"}
	moviePartitionKeys = []string{"movie_id"}
	movieSortKeys  []string
)

const (
	movieTableName = "movies"
)

func getTable(tableName string, cols, partitionKeys, sortKeys []string) db.Table{
	tableMeta := table.Metadata{
		Columns: cols,
		PartKey: partitionKeys,
		SortKey: sortKeys,
	}
	return table.New(tableMeta)
}

func insertMoviesInDb(ctx context.Context, session db.SessionWrapperService, movies *daomodels.Movies ) error{
	configTable := getTable(movieTableName, movieColumns, moviePartitionKeys, movieSortKeys)
	query := session.Query(configTable.Insert()).BindStruct(movies)
	if err := query.Exec(ctx); err != nil{
		return err
	}
	return nil
}