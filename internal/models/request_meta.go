package models

import "github.com/go-grpc-service/commons/constants"

type RequestMeta struct {
	uri       string
	movieName string
	desc      string
	genre     string
	ratings   string
}

func newRequestMeta(uri string, movieName, desc, genre, ratings string) *RequestMeta {
	return &RequestMeta{
		uri:       uri,
		movieName: movieName,
		desc:      desc,
		genre:     genre,
		ratings:   ratings,
	}
}
func CreateRequestMeta(uri, movieName, desc, genre, ratings string) *RequestMeta {
	return newRequestMeta(uri, movieName, desc, genre, ratings)
}

func CreateRequestMetaWithoutMovieDetails(uri string, movieName string) *RequestMeta {
	return newRequestMeta(uri, movieName, constants.EmptyString, constants.EmptyString, constants.EmptyString)
}

func InvalidRequestMeta() *RequestMeta {
	return newRequestMeta(constants.InvalidApi, constants.EmptyString, constants.EmptyString, constants.EmptyString, constants.EmptyString)
}

func (c *RequestMeta) GetUri() string {
	return c.uri
}

func (c *RequestMeta) GetMovieName() string {
	return c.movieName
}

func (c *RequestMeta) GetGenre() string {
	return c.genre
}

func (c *RequestMeta) GetDesc() string {
	return c.desc
}

func (c *RequestMeta) GetRatings() string {
	return c.ratings
}