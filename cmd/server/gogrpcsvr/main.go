package main

import (
	"context"
	"flag"
	"github.com/go-grpc-service/commons/constants"
	"github.com/go-grpc-service/internal/config"
	"github.com/go-grpc-service/internal/grpc"
	"github.com/go-grpc-service/internal/http"
	"github.com/go-grpc-service/internal/log"
	"github.com/go-grpc-service/pkg/grpcserver"
	"github.com/go-grpc-service/resources/moviepb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	netHttp "net/http"
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
	router := http.NewRouter()
	initHttpModules(router, cfg)
	movieServer, server := initGrpcModules(cfg)
	go runGatewayServer(cfg, movieServer)
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

func initHttpModules(router http.Router, cfg *config.AppConfig) {
	log.Logger.Info("adding static asset route")
	router.StaticRoute()
	server := http.NewServer(cfg, router)
	log.Logger.Info("HttpSever running")
	go server.Start()
}

func runGatewayServer(configurations *config.AppConfig, movieServer *grpcserver.MovieGrpcServer) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcMux := runtime.NewServeMux(
		runtime.WithMetadata(func(c context.Context, req *netHttp.Request) metadata.MD {
			return metadata.Pairs("x-forwarded-method", req.Method)
		}))
	opts := []grpc2.DialOption{grpc2.WithTransportCredentials(insecure.NewCredentials())}
	grpcServerEndpoint := flag.String(grpcServerEndpointName, constants.ColonSeparator+configurations.GRPCPort, grpcEndpointNameUsage)
	if err := moviepb.RegisterMoviePlatformHandlerFromEndpoint(ctx, grpcMux, *grpcServerEndpoint, opts); err != nil {
		log.Logger.Fatal("cannot register gateway handler server")
	}
	if err := netHttp.ListenAndServe(constants.ColonSeparator+configurations.ReverseProxyHttpPort, grpcMux); err != nil {
		log.Logger.Fatal("Failed to serve to", zap.String("port", configurations.ReverseProxyHttpPort), zap.Error(err))
	}
}
