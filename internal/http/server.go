package http

import (
	"github.com/go-grpc-service/internal/config"
	"github.com/go-grpc-service/internal/log"
	"github.com/gorilla/handlers"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type Server struct {
	*zap.Logger
	config *config.AppConfig
	srv    *http.Server
	mux    http.Handler
}

func NewServer(config *config.AppConfig, router Router) *Server {
	return &Server{
		config: config,
		mux:    router.Mux(),
		Logger: log.Logger,
	}
}

func (s *Server) Start() {
	srv := &http.Server{
		Handler:      handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(s.mux),
		ReadTimeout:  s.config.HttpReadTimeout,
		WriteTimeout: s.config.HttpWriteTimeout,
		IdleTimeout:  s.config.HttpIdleTimeout,
	}
	s.srv = srv

	err := srv.ListenAndServe()
	if err != nil && !strings.Contains(err.Error(), "http: Server Closed") {
		s.Logger.Fatal("failed to start server %v", zap.Error(err))
	}
}
