package interceptor

import (
	"context"
	"github.com/go-grpc-service/commons/constants"
	"github.com/go-grpc-service/commons/utils"
	"github.com/go-grpc-service/helper"
	"github.com/go-grpc-service/internal/models"
	"github.com/go-grpc-service/resources/moviepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"
	"time"
)

const (
	proto                  = "/MoviePlatform/"
	CreateMovieFullMethod  = proto + constants.CreateMovieApi
	UpdateMovieFullMethod  = proto + constants.UpdateMovieApi
	GetMovieFullMethod     = proto + constants.GetMovieApi
	GetAllMoviesFullMethod = proto + constants.GetAllMoviesApi
	HealthCheckURI         = "/grpc.health.v1.Health/Check"
	invalid                = "invalid"
	missing                = "missing"
	panic                  = "panic"
)

func RequestInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	meta := ConvertRequestToRequestMeta(req, info)
	requestValidationErr := ValidateRequest(req, info, meta)
	if requestValidationErr != nil {
		errStr := requestValidationErr.Error()
		statusCode := getErrorStatusCode(errStr)
		returnErr := status.Error(statusCode, requestValidationErr.Error())
		return nil, returnErr
	}
	defer panicRecovery(meta, start)
	ctx = context.WithValue(ctx, constants.RequestMeta, meta)

	response, err := handler(ctx, req)
	statusCode := codes.OK
	var returnErr error
	if err != nil {
		errStr := err.Error()
		statusCode = getErrorStatusCode(errStr)
		returnErr = status.Error(statusCode, err.Error())
	}

	log.Printf("Request completed for URI :  %s , MovieName : %s Genre :  %s , Desc : %s , TimeTaken(ms) :  %s, Error : %v",
		meta.GetUri(), meta.GetMovieName(), meta.GetGenre(), meta.GetDesc(), utils.GetTimeTakenInString(start), returnErr)
	return response, returnErr
}

func panicRecovery(meta *models.RequestMeta, startTime time.Time) {
	if err := recover(); err != nil {
		log.Printf("panic caught for URI :  %s , MovieName : %s Genre :  %s , Desc : %s , TimeTaken(ms) :  %s, Error : %v",
			meta.GetUri(), meta.GetMovieName(), meta.GetGenre(), meta.GetDesc(), utils.GetTimeTakenInString(startTime), err.(error).Error())
	}
}

func ValidateRequest(req interface{}, info *grpc.UnaryServerInfo, meta *models.RequestMeta) error {
	switch info.FullMethod {
	case CreateMovieFullMethod:
		r := req.(*moviepb.MovieRequest)
		return helper.ValidateCreateMovieRequest(r)
	case UpdateMovieFullMethod:
		r := req.(*moviepb.UpdateMovieRequest)
		return helper.ValidateUpdateMovieRequest(r)
	case GetMovieFullMethod:
		r := req.(*moviepb.GetMovieRequest)
		return helper.ValidateGetMovieRequest(r)
	case GetAllMoviesFullMethod:
		r := req.(*moviepb.GetAllMoviesRequest)
		return helper.ValidateGetAllMoviesRequest(r)
	}

	return nil
}

func ConvertRequestToRequestMeta(req interface{}, info *grpc.UnaryServerInfo) *models.RequestMeta {
	switch info.FullMethod {
	case CreateMovieFullMethod:
		r := req.(*moviepb.MovieRequest)
		return models.CreateRequestMeta(constants.CreateMovieApi, r.GetMovie(), r.GetGenre(), r.GetRating(), r.GetDesc())
	case UpdateMovieFullMethod:
		r := req.(*moviepb.UpdateMovieRequest)
		return models.CreateRequestMeta(constants.UpdateMovieApi, r.GetMovie(), r.GetGenre(), r.GetRating(), r.GetDesc())
	case GetMovieFullMethod:
		r := req.(*moviepb.GetMovieRequest)
		return models.CreateRequestMetaWithoutMovieDetails(constants.GetMovieApi, r.GetMovie())
	case GetAllMoviesFullMethod:
		r := req.(*moviepb.GetAllMoviesRequest)
		return models.CreateRequestMetaWithoutMovieDetails(constants.GetAllMoviesApi, r.GetMovies())
	}
	return models.InvalidRequestMeta()
}

func getErrorStatusCode(errStr string) codes.Code {
	code := codes.Internal
	if strings.Contains(errStr, invalid) || strings.Contains(errStr, missing) {
		code = codes.InvalidArgument
	}
	return code
}