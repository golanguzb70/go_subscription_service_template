package services

import (
	"context"

	pb "github.com/golanguzb70/go_subscription_service/genproto/subscription_service"
	l "github.com/golanguzb70/go_subscription_service/pkg/logger"
	grpclient "github.com/golanguzb70/go_subscription_service/server/grpc/client"
	"github.com/golanguzb70/go_subscription_service/storage"
)

type SubscriptionCategoryService struct {
	storage  storage.StorageI
	logger   l.Logger
	services grpclient.ServiceManager
	pb.UnimplementedSubscriptionCategoryServiceServer
}

// New Category Service
func NewSubscriptionCategoryService(stroge storage.StorageI, log l.Logger, services grpclient.ServiceManager) *SubscriptionCategoryService {
	return &SubscriptionCategoryService{
		storage:  stroge,
		logger:   log,
		services: services,
	}
}

func (s *SubscriptionCategoryService) Create(ctx context.Context, req *pb.SubscriptionCategory) (*pb.SubscriptionCategory, error) {
	return s.storage.SubscriptionCategory().Create(ctx, req)
}

func (s *SubscriptionCategoryService) Get(ctx context.Context, req *pb.Id) (*pb.SubscriptionCategory, error) {
	return s.storage.SubscriptionCategory().Get(ctx, req)
}

func (s *SubscriptionCategoryService) Find(ctx context.Context, req *pb.GetListFilter) (*pb.SubscriptionCategories, error) {
	return s.storage.SubscriptionCategory().Find(ctx, req)
}

func (s *SubscriptionCategoryService) Update(ctx context.Context, req *pb.SubscriptionCategory) (*pb.SubscriptionCategory, error) {
	return s.storage.SubscriptionCategory().Update(ctx, req)
}

func (s *SubscriptionCategoryService) Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	return s.storage.SubscriptionCategory().Delete(ctx, req)
}
