package dao

import (
	"context"
	"errors"
	"sync"

	"github.com/go-grpc-service/internal/config"
	daomodels "github.com/go-grpc-service/internal/dao/dao_models"
	"github.com/go-grpc-service/internal/db"
	"github.com/go-grpc-service/internal/models"
)

var (initSync sync.Once
	movieDao MovieDao
)

//change the req with a struct of request and also add a repsonse
type MovieDao interface{
	InsertMovie(ctx context.Context, req *models.Movie)(*models.Movie, error);

}

type MovieImpl struct {
	SessionWrapper db.SessionWrapperService
}

func NewMovieDaoImpl(dbConfigs config.DbConfigs) MovieDao{
	initSync.Do( func() {
		movieDao = &MovieImpl{
			SessionWrapper: db.GetSession(dbConfigs),
		}
	})
	return movieDao
}

func (dao *MovieImpl) InsertMovie(ctx context.Context, req *models.Movie)(*models.Movie, error) {

	daoMovie := convertToDaoMovie(req)
	if err := insertMoviesInDb(ctx, dao.SessionWrapper, daoMovie); err != nil {
		return nil, errors.New("error inserting key in db" + req.Name)
	}
	return req, nil
}

func convertToDaoMovie(req *models.Movie) *daomodels.Movies{

	return &daomodels.Movies{
		Name: req.Name,
		Genre:  req.Genre,
		Description:  req.Desc,
		Rating:  	req.Rating,
	}

}

// type ConfigDaoImpl struct{
// 	SessionWrapper 

// }

// func NewConfigDaoImpl()ConfigDao{
// 	initSync.Do(func(){


// 	})

// }