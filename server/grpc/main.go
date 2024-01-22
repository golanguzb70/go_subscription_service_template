package grpc

import (
	"fmt"

	"net"

	"github.com/golanguzb70/go_subscription_service/config"
	pb "github.com/golanguzb70/go_subscription_service/genproto/subscription_service"
	l "github.com/golanguzb70/go_subscription_service/pkg/logger"
	grpclient "github.com/golanguzb70/go_subscription_service/server/grpc/client"
	"github.com/golanguzb70/go_subscription_service/server/grpc/services"
	"github.com/golanguzb70/go_subscription_service/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

type GRPCService struct {
	ResourceCategoryService *services.ResourceCategoryService
}

func New(cfg *config.Config, logger l.Logger) (*GRPCService, error) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting to database: %v", err)
	}

	grpcClient, err := grpclient.NewGrpcClients(cfg)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting with grpc clients: %v", err)
	}

	return &GRPCService{
		ResourceCategoryService: services.NewResourceCategoryService(storage.New(connDb), logger, grpcClient),
	}, nil
}

func (service *GRPCService) Run(logger l.Logger, cfg *config.Config) {
	server := grpc.NewServer()

	pb.RegisterResourceCategoryServiceServer(server, service.ResourceCategoryService)

	listener, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		panic("Error while creating listener")
	}

	logger.Info("GRPC server is running on port " + cfg.RPCPort)

	err = server.Serve(listener)
	if err != nil {
		panic("error while serving gRPC server on port " + cfg.RPCPort)
	}
}
