package grpc

import (
	"fmt"
	"github.com/go-grpc-service/internal/grpc/interceptor"
	"log"
	"net"

	"github.com/go-grpc-service/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
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
	listener, err := net.Listen("tcp", ":"+s.Config.GRPCPort)
	if err != nil {
		log.Fatal(err)
	}
	s.HealthServer.SetServingStatus("go-grpc-service", healthv1.HealthCheckResponse_SERVING)
	fmt.Print("grpc server started")
	if err := s.GrpcServer.Serve(listener); err != nil {
		s.HealthServer.SetServingStatus("go-grpc-service", healthv1.HealthCheckResponse_NOT_SERVING)
		log.Fatal(err)
	}

}