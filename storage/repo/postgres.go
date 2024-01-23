package repo

import (
	"context"

	pb "github.com/golanguzb70/go_subscription_service/genproto/subscription_service"
)

type ResourceCategoryI interface {
	Create(ctx context.Context, req *pb.ResourceCategory) (*pb.ResourceCategory, error)
	Get(ctx context.Context, req *pb.Id) (*pb.ResourceCategory, error)
	Find(ctx context.Context, req *pb.GetListFilter) (*pb.ResourceCategories, error)
	Update(ctx context.Context, req *pb.ResourceCategory) (*pb.ResourceCategory, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error)
}
