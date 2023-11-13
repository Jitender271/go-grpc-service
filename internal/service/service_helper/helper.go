package service_helper

import (
	"context"
	"github.com/go-grpc-service/internal/dao"
)

func IsDuplicateMovie(ctx context.Context, movieDao dao.MovieDao, movieName string) (bool, error) {
	movieDetails, getMovieDetailsErr := movieDao.GetMovie(ctx, movieName)

	if movieDetails == nil {
		return false, nil
	}
	if getMovieDetailsErr != nil {
		return false, getMovieDetailsErr
	}
	return true, nil
}
