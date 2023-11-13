package service

import (
	"context"
	"errors"
	"github.com/go-grpc-service/internal/config"
	"github.com/go-grpc-service/internal/dao"
	daomodels "github.com/go-grpc-service/internal/dao/dao_models"
	"github.com/go-grpc-service/internal/models"
	"github.com/go-grpc-service/internal/service/service_helper"
	"github.com/go-grpc-service/resources/moviepb"
)

const (
	duplicateMovieError = "movie already exists"
	movieError          = "error validating movie details"
)

type MovieService interface {
	CreateMovie(ctx context.Context, request *moviepb.MovieRequest) (*models.Movie, error)
	GetMovie(ctx context.Context, request *moviepb.GetMovieRequest) (*models.Movie, error)
	GetAllMovies(ctx context.Context) ([]*models.Movie, error)
	UpdateMovies(ctx context.Context, request *moviepb.UpdateMovieRequest) (*models.Movie, error)
}

type MovieServiceImpl struct {
	movieDao dao.MovieDao
}

func NewMovieImpl(dbConfigs config.DbConfigs) MovieService {
	return &MovieServiceImpl{
		movieDao: dao.NewMovieDaoImpl(dbConfigs),
	}
}

func (m *MovieServiceImpl) CreateMovie(ctx context.Context, request *moviepb.MovieRequest) (*models.Movie, error) {
	isDuplicate, getMovieDetailsErr := service_helper.IsDuplicateMovie(ctx, m.movieDao, request.GetMovie())

	if getMovieDetailsErr != nil {
		return nil, errors.New(movieError)
	}

	if isDuplicate {
		return nil, errors.New(duplicateMovieError)
	}

	movie, err := m.movieDao.InsertMovie(ctx, getMovieModel(request))
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (m *MovieServiceImpl) GetMovie(ctx context.Context, request *moviepb.GetMovieRequest) (*models.Movie, error) {
	movie, err := m.movieDao.GetMovie(ctx, request.GetMovie())
	if err != nil {
		return nil, err
	}
	mov := getMovieModelFromDao(movie)
	return mov, nil
}

func (m *MovieServiceImpl) GetAllMovies(ctx context.Context) ([]*models.Movie, error) {
	movies, err := m.movieDao.GetAllMovies(ctx)
	if err != nil {
		return nil, err
	}
	var movieList []*models.Movie
	for _, movie := range movies {
		movieList = append(movieList, getMovieModelFromDao(&movie))
	}
	return movieList, nil
}

func (m *MovieServiceImpl) UpdateMovies(ctx context.Context, request *moviepb.UpdateMovieRequest) (*models.Movie, error) {
	_, err := validateMovieExists(ctx, m, request)
	if err != nil {
		return nil, err
	}
	movie, err := m.movieDao.UpdateMovies(ctx, getUpdatedMovieModel(request))
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func getMovieModel(request *moviepb.MovieRequest) *models.Movie {
	return &models.Movie{
		Name:   request.GetMovie(),
		Desc:   request.GetDesc(),
		Genre:  request.GetGenre(),
		Rating: request.GetRating(),
	}
}

func getUpdatedMovieModel(request *moviepb.UpdateMovieRequest) *models.Movie {
	return &models.Movie{
		Id:     request.GetMovieId(),
		Name:   request.GetMovie(),
		Desc:   request.GetDesc(),
		Genre:  request.GetGenre(),
		Rating: request.GetRating(),
	}
}

func getMovieModelFromDao(request *daomodels.Movies) *models.Movie {
	return &models.Movie{
		Id:     request.MovieID,
		Name:   request.Name,
		Desc:   request.Description,
		Genre:  request.Genre,
		Rating: request.Rating,
	}
}

func validateMovieExists(ctx context.Context, m *MovieServiceImpl, request *moviepb.UpdateMovieRequest) (*models.Movie, error) {
	movieDetails, err := getMovieDetails(ctx, m, request.GetMovie())

	if err != nil {
		return nil, err
	}
	if movieDetails == nil || movieDetails.MovieID != request.GetMovieId() {
		return nil, errors.New("movie name :" + request.Movie + ", id : " + request.MovieId + "combination doesn't match any record")
	}
	return getMovieModelFromDao(movieDetails), nil
}

func getMovieDetails(ctx context.Context, m *MovieServiceImpl, movieName string) (*daomodels.Movies, error) {
	movie, err := m.movieDao.GetMovie(ctx, movieName)
	if err != nil {
		return nil, err
	}
	if movie == nil {
		return nil, nil
	}
	return movie, nil
}
