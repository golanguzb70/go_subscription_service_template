package repo

import (
	"context"

	pb "github.com/golanguzb70/go_subscription_service/genproto/subscription_service"
	"github.com/golanguzb70/go_subscription_service/storage/models"
)

type ResourceCategoryI interface {
	Create(ctx context.Context, req *pb.ResourceCategory) (*pb.ResourceCategory, error)
	Get(ctx context.Context, req *pb.Id) (*pb.ResourceCategory, error)
	Find(ctx context.Context, req *pb.GetListFilter) (*pb.ResourceCategories, error)
	Update(ctx context.Context, req *pb.ResourceCategory) (*pb.ResourceCategory, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error)
	AddResource(ctx context.Context, req *pb.ResourceAndCategoryIds) (*pb.Empty, error)
	RemoveResource(ctx context.Context, req *pb.ResourceAndCategoryIds) (*pb.Empty, error)
}

type ResourceI interface {
	Create(ctx context.Context, req *pb.Resource) (*pb.Resource, error)
	Get(ctx context.Context, req *pb.Id) (*pb.Resource, error)
	Find(ctx context.Context, req *pb.GetListFilter) (*pb.Resources, error)
	Update(ctx context.Context, req *pb.Resource) (*pb.Resource, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error)
}

type SubscriptionCategoryI interface {
	Create(ctx context.Context, req *pb.SubscriptionCategory) (*pb.SubscriptionCategory, error)
	Get(ctx context.Context, req *pb.Id) (*pb.SubscriptionCategory, error)
	Find(ctx context.Context, req *pb.GetListFilter) (*pb.SubscriptionCategories, error)
	Update(ctx context.Context, req *pb.SubscriptionCategory) (*pb.SubscriptionCategory, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error)
	AddResourceCategory(ctx context.Context, req *pb.SubscriptionResourceCategoryIds) (*pb.Empty, error)
	RemoveResourceCategory(ctx context.Context, req *pb.SubscriptionResourceCategoryIds) (*pb.Empty, error)
}

type SubscriptionI interface {
	Create(ctx context.Context, req *pb.Subscription) (*pb.Subscription, error)
	Get(ctx context.Context, req *pb.Id) (*pb.Subscription, error)
	Find(ctx context.Context, req *pb.GetListFilter) (*pb.Subscriptions, error)
	Update(ctx context.Context, req *pb.Subscription) (*pb.Subscription, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error)
}

type UserSubscriptionI interface {
	Buy(ctx context.Context, req *models.CreateUserSubscriptionReq) (*pb.Empty, error)
}
