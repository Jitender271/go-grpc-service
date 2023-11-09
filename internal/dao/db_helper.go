package dao

import (
	"context"
	"strings"

	daomodels "github.com/go-grpc-service/internal/dao/dao_models"
	"github.com/go-grpc-service/internal/db"
	"github.com/scylladb/gocqlx/v2/table"
)

var (
	movieColumns       = []string{"movie_id", "name", "genre", "description", "rating"}
	moviePartitionKeys = []string{"name"}
	updateMovieColumns = []string{"genre", "description", "rating"}
	movieSortKeys      []string
)

const (
	movieTableName = "movies"
	name           = "name"
)

func getTable(tableName string, cols, partitionKeys, sortKeys []string) db.Table {
	tableMeta := table.Metadata{
		Name:    tableName,
		Columns: cols,
		PartKey: partitionKeys,
		SortKey: sortKeys,
	}
	return table.New(tableMeta)
}

func insertMoviesInDb(ctx context.Context, session db.SessionWrapperService, movies *daomodels.Movies) error {
	configTable := getTable(movieTableName, movieColumns, moviePartitionKeys, movieSortKeys)
	query := session.Query(configTable.Insert()).BindStruct(movies)
	if err := query.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func getMovieFromDb(ctx context.Context, session db.SessionWrapperService, movieName string) (*daomodels.Movies, error) {
	var movies daomodels.Movies
	configTable := getTable(movieTableName, movieColumns, moviePartitionKeys, movieSortKeys)
	query := session.Query(configTable.Get()).BindMap(map[string]interface{}{name: movieName})
	err := query.GetRelease(ctx, &movies)
	if err != nil && strings.EqualFold(err.Error(), "not found") {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return &movies, nil
}

func getAllMoviesFromDb(ctx context.Context, session db.SessionWrapperService) ([]daomodels.Movies, error) {
	var movies []daomodels.Movies
	configTable := getTable(movieTableName, movieColumns, moviePartitionKeys, movieSortKeys)
	query := session.Query(configTable.SelectAll())
	err := query.SelectRelease(ctx, &movies)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func updateMoviesInDb(ctx context.Context, session db.SessionWrapperService, movies *daomodels.Movies) error {
	configTable := getTable(movieTableName, movieColumns, moviePartitionKeys, movieSortKeys)
	query := session.Query(configTable.Update(updateMovieColumns...)).BindStruct(movies)
	if err := query.Exec(ctx); err != nil {
		return err
	}
	return nil
}