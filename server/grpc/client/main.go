package grpclient

import (
	"github.com/golanguzb70/go_subscription_service/config"
)

type ServiceManager interface {
}

type grpcClients struct {
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	// connect to external clients here
	return &grpcClients{}, nil
}
