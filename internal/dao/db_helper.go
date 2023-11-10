package dao

import (
	"context"
	"github.com/go-grpc-service/internal/log"
	"go.uber.org/zap"
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
	id             = "id"
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
		log.Logger.Error("error inserting movie name in db ", zap.String("key", movies.Name))
		return err
	}
	log.Logger.Info("Movie insertion successful", zap.String(id, movies.MovieID))
	return nil
}

func getMovieFromDb(ctx context.Context, session db.SessionWrapperService, movieName string) (*daomodels.Movies, error) {
	var movies daomodels.Movies
	configTable := getTable(movieTableName, movieColumns, moviePartitionKeys, movieSortKeys)
	query := session.Query(configTable.Get()).BindMap(map[string]interface{}{name: movieName})
	err := query.GetRelease(ctx, &movies)
	if err != nil && strings.EqualFold(err.Error(), "not found") {
		log.Logger.Error("movie name not found in db ", zap.String("movie_name", movieName))
		return nil, err
	} else if err != nil {
		log.Logger.Error("error fetching movie details from db ", zap.Error(err))
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
		log.Logger.Error("error fetching all movie details from db ", zap.Error(err))
		return nil, err
	}
	return movies, nil
}

func updateMoviesInDb(ctx context.Context, session db.SessionWrapperService, movies *daomodels.Movies) error {
	configTable := getTable(movieTableName, movieColumns, moviePartitionKeys, movieSortKeys)
	query := session.Query(configTable.Update(updateMovieColumns...)).BindStruct(movies)
	if err := query.Exec(ctx); err != nil {
		log.Logger.Error("error updating movie", zap.String("key", movies.Name))
		return err
	}
	log.Logger.Info("Movie updated successful", zap.String(id, movies.MovieID))
	return nil
}
