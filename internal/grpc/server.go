package grpc

import (
	"github.com/go-grpc-service/internal/config"
	"github.com/go-grpc-service/internal/grpc/interceptor"
	log2 "github.com/go-grpc-service/internal/log"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

const (
	tcp               = "tcp"
	grpcHealthService = "health-service-grpc"
	moviePlatform     = "movie-platform"
	port              = "port"
)

type Server struct {
	Config       *config.AppConfig
	GrpcServer   *grpc.Server
	HealthServer *health.Server
}

func NewServer(config *config.AppConfig) *Server {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptor.RequestInterceptor),
		grpc.ConnectionTimeout(config.GRPCConnectionTimeout),
	)
	healthServer := health.NewServer()
	healthServer.SetServingStatus("health-grpc-service", healthv1.HealthCheckResponse_SERVING)
	healthv1.RegisterHealthServer(grpcServer, healthServer)
	return &Server{
		Config:       config,
		GrpcServer:   grpcServer,
		HealthServer: healthServer,
	}
}

func (s *Server) Start() {
	listener, err := net.Listen(tcp, ":"+s.Config.GRPCPort)
	if err != nil {
		log2.Logger.Fatal("failed to listen to", zap.String(port, s.Config.GRPCPort), zap.Error(err))
	}
	s.HealthServer.SetServingStatus("go-grpc-service", healthv1.HealthCheckResponse_SERVING)
	log2.Logger.Info("grpc server started")
	if err := s.GrpcServer.Serve(listener); err != nil {
		s.HealthServer.SetServingStatus("go-grpc-service", healthv1.HealthCheckResponse_NOT_SERVING)
		log2.Logger.Fatal("failed to serve", zap.Error(err))
	}

}
