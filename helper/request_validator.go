package helper

import (
	"errors"
	"github.com/go-grpc-service/commons/constants"
	"github.com/go-grpc-service/resources/moviepb"
)

func ValidateGetMovieRequest(req *moviepb.GetMovieRequest) error {
	if req == nil {
		return errors.New("nil struct received for validation")
	}
	if req.GetMovie() == constants.EmptyString {
		return errors.New("invalid request - missing tenant name")
	}
	return nil
}

func ValidateCreateMovieRequest(req *moviepb.MovieRequest) error {
	if req == nil {
		return errors.New("nil struct received for validation")
	}
	if req.GetMovie() == constants.EmptyString {
		return errors.New("invalid request - missing movie name")
	}
	if req.GetDesc() == constants.EmptyString {
		return errors.New("invalid request - missing movie description")
	}
	if req.GetGenre() == constants.EmptyString {
		return errors.New("invalid request - missing movie genre")
	}
	if req.GetRating() == constants.EmptyString {
		return errors.New("invalid request - missing movie ratings")
	}
	return nil
}

func ValidateUpdateMovieRequest(req *moviepb.UpdateMovieRequest) error {
	if req == nil {
		return errors.New("nil struct received for validation")
	}
	if req.GetMovie() == constants.EmptyString {
		return errors.New("invalid request - missing movie name")
	}
	if req.GetDesc() == constants.EmptyString {
		return errors.New("invalid request - missing movie description")
	}
	if req.GetGenre() == constants.EmptyString {
		return errors.New("invalid request - missing movie genre")
	}
	if req.GetRating() == constants.EmptyString {
		return errors.New("invalid request - missing movie ratings")
	}
	return nil
}

func ValidateGetAllMoviesRequest(req *moviepb.GetAllMoviesRequest) error {
	if req == nil {
		return errors.New("nil struct received for validation")
	}
	return nil
}