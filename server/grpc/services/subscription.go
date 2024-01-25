package services

import (
	"context"

	pb "github.com/golanguzb70/go_subscription_service/genproto/subscription_service"
	l "github.com/golanguzb70/go_subscription_service/pkg/logger"
	grpclient "github.com/golanguzb70/go_subscription_service/server/grpc/client"
	"github.com/golanguzb70/go_subscription_service/storage"
)

type SubscriptionService struct {
	storage  storage.StorageI
	logger   l.Logger
	services grpclient.ServiceManager
	pb.UnimplementedSubscriptionServiceServer
}

// New Category Service
func NewSubscriptionService(stroge storage.StorageI, log l.Logger, services grpclient.ServiceManager) *SubscriptionService {
	return &SubscriptionService{
		storage:  stroge,
		logger:   log,
		services: services,
	}
}

func (s *SubscriptionService) Create(ctx context.Context, req *pb.Subscription) (*pb.Subscription, error) {
	return s.storage.Subscription().Create(ctx, req)
}

func (s *SubscriptionService) Get(ctx context.Context, req *pb.Id) (*pb.Subscription, error) {
	return s.storage.Subscription().Get(ctx, req)
}

func (s *SubscriptionService) Find(ctx context.Context, req *pb.GetListFilter) (*pb.Subscriptions, error) {
	return s.storage.Subscription().Find(ctx, req)
}

func (s *SubscriptionService) Update(ctx context.Context, req *pb.Subscription) (*pb.Subscription, error) {
	return s.storage.Subscription().Update(ctx, req)
}

func (s *SubscriptionService) Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	return s.storage.Subscription().Delete(ctx, req)
}
