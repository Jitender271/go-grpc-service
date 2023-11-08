package service

import (
	"context"

	"fmt"

	"github.com/go-grpc-service/internal/config"
	"github.com/go-grpc-service/internal/dao"
	"github.com/go-grpc-service/internal/models"
	"github.com/go-grpc-service/resources/moviepb"
)

type MovieService interface{
	CreateMovie(ctx context.Context, request *moviepb.MovieRequest)(*models.Movie, error)
}

type MovieServiceImpl struct{
	movieDao dao.MovieDao
}

func NewMovieImpl(dbConfigs config.DbConfigs) MovieService{
	return  &MovieServiceImpl{
		movieDao: dao.NewMovieDaoImpl(dbConfigs),
	}
}

func (m *MovieServiceImpl) CreateMovie(ctx context.Context, request *moviepb.MovieRequest)(*models.Movie, error){
	fmt.Print("reaching")
	movie , err := m.movieDao.InsertMovie(ctx, getMovieModel(request))
	if err != nil {
		return nil, err
	}
	fmt.Print("movie", movie)
	return movie, nil

}

func getMovieModel(request *moviepb.MovieRequest) *models.Movie{
	return &models.Movie{
		Name: request.GetMovie(),
		Desc: request.GetDesc(),
		Genre: request.GetGenre(),
		Rating: request.GetRating(),
	}

}