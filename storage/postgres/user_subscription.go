package postgres

import (
	"context"

	pb "github.com/golanguzb70/go_subscription_service/genproto/subscription_service"
	"github.com/golanguzb70/go_subscription_service/pkg/db"
	"github.com/golanguzb70/go_subscription_service/pkg/logger"
	"github.com/golanguzb70/go_subscription_service/storage/models"
	"github.com/golanguzb70/go_subscription_service/storage/repo"
	"github.com/google/uuid"
)

type userSubscription struct {
	db  *db.Postgres
	log logger.Logger
}

func NewUserSubscriptionRepo(db *db.Postgres, log logger.Logger) repo.UserSubscriptionI {
	return &userSubscription{
		db:  db,
		log: log,
	}
}

func (r *userSubscription) Buy(ctx context.Context, req *models.CreateUserSubscriptionReq) (*pb.Empty, error) {
	query := r.db.Builder.Insert("user_subscriptions").
		Columns(`id, user_key, subscription_id, start_time, end_time, active`).
		Values(uuid.New().String(), req.UserId, req.SubscriptionId, req.StartTime, req.EndTime, req.Active)

	_, err := query.RunWith(r.db.Db).Exec()
	if err != nil {
		return nil, HandleDatabaseError(err, r.log, "create resource")
	}

	return &pb.Empty{}, nil
}

func (r *userSubscription) CheckSubscription(ctx context.Context, req *pb.CheckSubscriptionRequest) (*pb.CheckSubscriptionResponse, error) {
	
	return nil, nil
}
