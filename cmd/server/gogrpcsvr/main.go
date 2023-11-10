package main

import (
	"github.com/go-grpc-service/internal/config"
	"github.com/go-grpc-service/internal/grpc"
	log "github.com/go-grpc-service/internal/log"
	"github.com/go-grpc-service/pkg/grpcserver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logLevelDebug          = "DEBUG"
	grpcServerEndpointName = "grpc-server-endpoint"
	grpcEndpointNameUsage  = "gRPC server endpoint"
)

func main() {
	cfg := config.InitConfig(false)
	initLogger(cfg)
	defer func() {
		if r := recover(); r != nil {
			log.Logger.Fatal("Error while starting app")
		}
	}()
	_, server := initGrpcModules(cfg)
	server.Start()
	log.Logger.Info("Application Started successfully on port 80")
}

func initGrpcModules(configuration *config.AppConfig) (*grpcserver.MovieGrpcServer, *grpc.Server) {
	log.Logger.Info("starting grpc server ...")
	server := grpc.NewServer(configuration)
	movieServer := grpcserver.NewGrpcServer(server)
	return movieServer, server
}

func initLogger(config *config.AppConfig) {
	level := zap.InfoLevel
	if config.LogLevel == logLevelDebug {
		level = zapcore.DebugLevel
	}
	log.NewLogger(level)
}
