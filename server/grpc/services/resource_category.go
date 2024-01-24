package services

import (
	"context"

	pb "github.com/golanguzb70/go_subscription_service/genproto/subscription_service"
	l "github.com/golanguzb70/go_subscription_service/pkg/logger"
	grpclient "github.com/golanguzb70/go_subscription_service/server/grpc/client"
	"github.com/golanguzb70/go_subscription_service/storage"
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

func (s *ResourceCategoryService) Create(ctx context.Context, req *pb.ResourceCategory) (*pb.ResourceCategory, error) {
	return s.storage.ResourceCategory().Create(ctx, req)
}

func (s *ResourceCategoryService) Get(ctx context.Context, req *pb.Id) (*pb.ResourceCategory, error) {
	return s.storage.ResourceCategory().Get(ctx, req)
}

func (s *ResourceCategoryService) Find(ctx context.Context, req *pb.GetListFilter) (*pb.ResourceCategories, error) {
	return s.storage.ResourceCategory().Find(ctx, req)
}

func (s *ResourceCategoryService) Update(ctx context.Context, req *pb.ResourceCategory) (*pb.ResourceCategory, error) {
	return s.storage.ResourceCategory().Update(ctx, req)
}

func (s *ResourceCategoryService) Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	return s.storage.ResourceCategory().Delete(ctx, req)
}

func (s *ResourceCategoryService) AddResource(ctx context.Context, req *pb.ResourceAndCategoryIds) (*pb.Empty, error) {
	return s.storage.ResourceCategory().AddResource(ctx, req)
}

func (s *ResourceCategoryService) RemoveResource(ctx context.Context, req *pb.ResourceAndCategoryIds) (*pb.Empty, error) {
	return s.storage.ResourceCategory().RemoveResource(ctx, req)
}
