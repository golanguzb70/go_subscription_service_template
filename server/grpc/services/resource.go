package services

import (
	"context"

	pb "github.com/golanguzb70/go_subscription_service/genproto/subscription_service"
	l "github.com/golanguzb70/go_subscription_service/pkg/logger"
	grpclient "github.com/golanguzb70/go_subscription_service/server/grpc/client"
	"github.com/golanguzb70/go_subscription_service/storage"
)

type ResourceService struct {
	storage  storage.StorageI
	logger   l.Logger
	services grpclient.ServiceManager
	pb.UnimplementedResourceServiceServer
}

// New Category Service
func NewResourceService(stroge storage.StorageI, log l.Logger, services grpclient.ServiceManager) *ResourceService {
	return &ResourceService{
		storage:  stroge,
		logger:   log,
		services: services,
	}
}

func (s *ResourceService) Create(ctx context.Context, req *pb.Resource) (*pb.Resource, error) {
	return s.storage.Resource().Create(ctx, req)
}

func (s *ResourceService) Get(ctx context.Context, req *pb.Id) (*pb.Resource, error) {
	return s.storage.Resource().Get(ctx, req)
}

func (s *ResourceService) Find(ctx context.Context, req *pb.GetListFilter) (*pb.Resources, error) {
	return s.storage.Resource().Find(ctx, req)
}

func (s *ResourceService) Update(ctx context.Context, req *pb.Resource) (*pb.Resource, error) {
	return s.storage.Resource().Update(ctx, req)
}

func (s *ResourceService) Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	return s.storage.Resource().Delete(ctx, req)
}
