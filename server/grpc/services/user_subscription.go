package services

import (
	"context"
	"time"

	pb "github.com/golanguzb70/go_subscription_service/genproto/subscription_service"
	l "github.com/golanguzb70/go_subscription_service/pkg/logger"
	grpclient "github.com/golanguzb70/go_subscription_service/server/grpc/client"
	"github.com/golanguzb70/go_subscription_service/storage"
	"github.com/golanguzb70/go_subscription_service/storage/models"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserSubscriptionService struct {
	storage  storage.StorageI
	logger   l.Logger
	services grpclient.ServiceManager
	pb.UnimplementedUserSubscriptionServiceServer
}

// New Category Service
func NewUserSubscriptionService(stroge storage.StorageI, log l.Logger, services grpclient.ServiceManager) *UserSubscriptionService {
	return &UserSubscriptionService{
		storage:  stroge,
		logger:   log,
		services: services,
	}
}

/*
Logic of buying new subscription:

1. Get the subscription & Check if it is active
2. Withdraw the amount from user's pocket
3. Activate subscription to the user
*/
func (s *UserSubscriptionService) Buy(ctx context.Context, req *pb.BuyRequest) (*pb.Empty, error) {
	//1. Get the subscription & Check if it is active
	subscription, err := s.storage.Subscription().Get(ctx, &pb.Id{
		Id: req.SubscriptionId,
	})
	if err != nil {
		return &pb.Empty{}, err
	}

	if !subscription.Active {
		return &pb.Empty{}, status.Error(codes.Unavailable, "this subscription is inactive")
	}

	// 2. Withdraw the amount from user's pocket

	// 3. Activate subscription to the user
	currentTime := time.Now()
	endTime := currentTime.AddDate(0, 0, int(subscription.Duration))

	_, err = s.storage.UserSubscription().Buy(ctx, &models.CreateUserSubscriptionReq{
		Id:             uuid.New().String(),
		UserId:         req.UserId,
		SubscriptionId: req.SubscriptionId,
		StartTime:      currentTime,
		EndTime:        endTime,
		Active:         true,
	})

	return &pb.Empty{}, err
}
