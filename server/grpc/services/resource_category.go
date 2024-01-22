package services

import (
	l "github.com/golanguzb70/go_subscription_service/pkg/logger"
	grpclient "github.com/golanguzb70/go_subscription_service/server/grpc/client"
	"github.com/golanguzb70/go_subscription_service/storage"
	pb "github.com/golanguzb70/go_subscription_service/genproto/subscription_service"
)

type ResourceCategoryService struct {
	storage  storage.StorageI
	logger   l.Logger
	services grpclient.ServiceManager
	pb.UnimplementedResourceCategoryServiceServer
}

// New Category Service
func NewResourceCategoryService(stroge storage.StorageI, log l.Logger, services grpclient.ServiceManager) *ResourceCategoryService {
	return &ResourceCategoryService{
		storage:  stroge,
		logger:   log,
		services: services,
	}
}
