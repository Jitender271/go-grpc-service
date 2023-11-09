package dao

import (
	"context"
	"errors"
	"github.com/gocql/gocql"
	"sync"

	"github.com/go-grpc-service/internal/config"
	daomodels "github.com/go-grpc-service/internal/dao/dao_models"
	"github.com/go-grpc-service/internal/db"
	"github.com/go-grpc-service/internal/models"
)

var (
	initSync sync.Once
	movieDao MovieDao
)

type MovieDao interface {
	InsertMovie(ctx context.Context, req *models.Movie) (*models.Movie, error)
	GetMovie(ctx context.Context, movieName string) (*daomodels.Movies, error)
	GetAllMovies(ctx context.Context) ([]daomodels.Movies, error)
	UpdateMovies(ctx context.Context, req *models.Movie) (*models.Movie, error)
}

type MovieImpl struct {
	SessionWrapper db.SessionWrapperService
}

func NewMovieDaoImpl(dbConfigs config.DbConfigs) MovieDao {
	initSync.Do(func() {
		movieDao = &MovieImpl{
			SessionWrapper: db.GetSession(dbConfigs),
		}
	})
	return movieDao
}

func (dao *MovieImpl) InsertMovie(ctx context.Context, req *models.Movie) (*models.Movie, error) {
	req.Id = gocql.TimeUUID().String()
	daoMovie := convertToDaoMovie(req)
	if err := insertMoviesInDb(ctx, dao.SessionWrapper, daoMovie); err != nil {
		return nil, errors.New("error inserting key in db" + req.Name)
	}
	return req, nil
}

func (dao *MovieImpl) GetMovie(ctx context.Context, movieName string) (*daomodels.Movies, error) {
	movie, err := getMovieFromDb(ctx, dao.SessionWrapper, movieName)
	if err != nil {
		return nil, errors.New("error getting movie from db" + movieName)
	}
	return movie, nil
}

func (dao *MovieImpl) GetAllMovies(ctx context.Context) ([]daomodels.Movies, error) {
	movies, err := getAllMoviesFromDb(ctx, dao.SessionWrapper)
	if err != nil {
		return nil, errors.New("error getting all movies from db")
	}
	return movies, nil
}

func (dao *MovieImpl) UpdateMovies(ctx context.Context, req *models.Movie) (*models.Movie, error) {
	daoMovie := convertToDaoMovie(req)
	if err := updateMoviesInDb(ctx, dao.SessionWrapper, daoMovie); err != nil {
		return nil, errors.New("error updating key in db" + req.Name)
	}
	return req, nil
}

func convertToDaoMovie(req *models.Movie) *daomodels.Movies {
	return &daomodels.Movies{
		MovieID:     req.Id,
		Name:        req.Name,
		Genre:       req.Genre,
		Description: req.Desc,
		Rating:      req.Rating,
	}

}