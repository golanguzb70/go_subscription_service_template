package main

import (
	"github.com/golanguzb70/go_subscription_service/config"
	"github.com/golanguzb70/go_subscription_service/pkg/logger"
	"github.com/golanguzb70/go_subscription_service/server/grpc"
)

func main() {
	cfg := config.Load()
	grpclog := logger.New(cfg.Environment, "go_subscription_service_grpc")

	services, err := grpc.New(cfg, grpclog)
	if err != nil {
		grpclog.Error("Error while initializing services", logger.Error(err))
		return
	}

	services.Run(grpclog, cfg)
}
