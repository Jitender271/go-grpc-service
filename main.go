package main

import (
	"github.com/go-grpc-service/internal/config"
	"github.com/go-grpc-service/internal/grpc"
	"github.com/go-grpc-service/pkg/grpcserver"
	"log"
)

func main() {
	cfg := config.InitConfig(false)
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("Error while starting app")
		}
	}()
	_, server := initGrpcModules(cfg)
	server.Start()
}

func initGrpcModules(configuration *config.AppConfig) (*grpcserver.MovieGrpcServer, *grpc.Server) {
	server := grpc.NewServer(configuration)
	movieServer := grpcserver.NewGrpcServer(server)
	return movieServer, server
}
