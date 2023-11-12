package db

import (
	configs "github.com/go-grpc-service/internal/config"
	"github.com/go-grpc-service/internal/log"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"go.uber.org/zap"
	"sync"
)

var (
	once                  sync.Once
	sessionWrapperService SessionWrapperService
)

type SessionWrapperService interface {
	Query(stmt string, names []string) QueryXService
}

type SessionWrapperImpl struct {
	Session *gocqlx.Session
}

func (s *SessionWrapperImpl) Query(stmt string, names []string) QueryXService {
	queryX := s.Session.Query(stmt, names)
	return &QueryXServiceImpl{
		QueryX: queryX,
	}
}

func GetSession(configs configs.DbConfigs) SessionWrapperService {
	once.Do(func() {
		scyllaDbCluster := gocql.NewCluster(configs.DBHosts...)
		scyllaDbCluster.Keyspace = configs.DBKeyspace
		scyllaDbCluster.ConnectTimeout = configs.DBConnectTimeout
		scyllaDbCluster.WriteTimeout = configs.DBWriteTimeout
		scyllaDbCluster.Timeout = configs.DBReadTimeout
		scyllaDbCluster.NumConns = configs.DBConnectionsPerHost
		scyllaDbCluster.Authenticator = gocql.PasswordAuthenticator{
			Username: configs.DBUsername,
			Password: configs.DBPassword,
		}
		s, err := gocqlx.WrapSession(scyllaDbCluster.CreateSession())
		if err != nil {
			log.Logger.Fatal("failed to create scylla db session ", zap.Error(err))
		}
		sessionWrapperService = &SessionWrapperImpl{
			Session: &s,
		}
	})
	return sessionWrapperService
}
