package service

import (
	"context"
	"github.com/go-grpc-service/internal/config"
	"github.com/go-grpc-service/internal/dao"
	daomodels "github.com/go-grpc-service/internal/dao/dao_models"
	"github.com/go-grpc-service/internal/models"
	"github.com/go-grpc-service/resources/moviepb"
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