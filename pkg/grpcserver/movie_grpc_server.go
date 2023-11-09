package grpcserver

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-grpc-service/internal/config"
	"github.com/go-grpc-service/internal/grpc"
	"github.com/go-grpc-service/internal/models"
	"github.com/go-grpc-service/internal/service"
	"github.com/go-grpc-service/resources/moviepb"
)

type MovieGrpcServer struct {
	moviepb.UnimplementedMoviePlatformServer
	service.MovieService
}

func NewGrpcServer(server *grpc.Server) *MovieGrpcServer {
	dbConfig := config.Configurations.DbConfigs
	s := &MovieGrpcServer{
		MovieService: service.NewMovieImpl(dbConfig),
	}
	moviepb.RegisterMoviePlatformServer(server.GrpcServer, s)
	return s
}

func (m *MovieGrpcServer) CreateMovie(ctx context.Context, request *moviepb.MovieRequest) (*moviepb.MovieResponse, error) {
	fmt.Print("ccp_grpc_server")
	movie, err := m.MovieService.CreateMovie(ctx, request)
	if err != nil {
		return nil, err
	}
	if movie == nil {
		return nil, errors.New("err not captured but nil response recieved")
	}

	return &moviepb.MovieResponse{
		MovieDetails: buildMovie(movie),
		Status:       moviepb.MovieStatus_CREATED,
	}, nil
}

func (m *MovieGrpcServer) GetMovie(ctx context.Context, request *moviepb.GetMovieRequest) (*moviepb.GetMovieResponse, error) {
	movie, err := m.MovieService.GetMovie(ctx, request)
	if err != nil {
		return nil, err
	}
	if movie == nil {
		return nil, errors.New("err not captured but nil response recieved")
	}

	return &moviepb.GetMovieResponse{
		MovieDetails: buildMovie(movie),
	}, nil
}

func (m *MovieGrpcServer) GetAllMovies(ctx context.Context, request *moviepb.GetAllMoviesRequest) (*moviepb.GetAllMoviesResponse, error) {
	movies, err := m.MovieService.GetAllMovies(ctx)
	if err != nil {
		return nil, err
	}
	if movies == nil {
		return nil, errors.New("err not captured but nil response recieved")
	}

	var movieList []*moviepb.MovieDetails
	for _, movie := range movies {
		movieList = append(movieList, buildMovie(movie))
	}
	return &moviepb.GetAllMoviesResponse{
		MovieDetails: movieList,
	}, nil
}

func (m *MovieGrpcServer) UpdateMovie(ctx context.Context, request *moviepb.UpdateMovieRequest) (*moviepb.UpdateMovieResponse, error) {
	movie, err := m.MovieService.UpdateMovies(ctx, request)
	if err != nil {
		return nil, err
	}
	if movie == nil {
		return nil, errors.New("err not captured but nil response received")
	}

	return &moviepb.UpdateMovieResponse{
		MovieDetails: buildMovie(movie),
		Status:       moviepb.MovieStatus_UPDATED,
	}, nil
}

func buildMovie(movie *models.Movie) *moviepb.MovieDetails {
	return &moviepb.MovieDetails{
		Id:          movie.Id,
		MovieName:   movie.Name,
		Genre:       movie.Genre,
		Description: movie.Desc,
		Ratings:     movie.Rating,
	}
}