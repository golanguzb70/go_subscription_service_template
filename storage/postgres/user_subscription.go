package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
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

func (r *userSubscription) CreateTvodAccess(ctx context.Context, req *pb.TvodAccess) (*pb.TvodAccess, error) {
	query := r.db.Builder.Insert("tvodacces").
		Columns(`id, user_key, resource_key, price`).
		Values(req.Id, req.UserId, req.ResourceKey, req.Price).
		Suffix("RETURNING start_time, created_at, updated_at")
	err := query.RunWith(r.db.Db).QueryRow().
		Scan(&req.StartTime, &req.CreatedAt, &req.UpdatedAt)

	return req, HandleDatabaseError(err, r.log, "create tvod access")
}

func (r *userSubscription) RemoveTvodAccess(ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	query := r.db.Builder.Delete("tvodacces").Where(squirrel.Eq{"id": req.Id})
	_, err := query.RunWith(r.db.Db).Exec()

	return &pb.Empty{}, HandleDatabaseError(err, r.log, "delete from tvodaccess")
}

func (r *userSubscription) CheckSubscription(ctx context.Context, req *pb.CheckSubscriptionRequest) (*pb.CheckSubscriptionResponse, error) {
	response := &pb.CheckSubscriptionResponse{}

	if req.Type == "svod" {
		query := `SELECT EXISTS(
			SELECT 1 FROM user_subscriptions us
			JOIN subscriptions s ON us.subscription_id = s.id
			JOIN subscription_categories sc ON s.category_id = sc.id
			JOIN resource_subsription_categories rsc ON sc.id=rsc.subscription_category_id
			JOIN resource_categories rc ON rc.id=rsc.resource_category_id 
			LEFT JOIN resources_categories_m2m rcm2m ON rsc.resource_category_id = rcm2m.category_id
			LEFT JOIN resources r ON rcm2m.resource_id=r.id  
			WHERE us.user_key=$1 AND us.active = true AND
			((rc.allow_all_resources = true AND rc.category_key = $2) OR (r.resource_key = $3 AND rc.allow_all_resources = false))
			);`

		err := r.db.Db.QueryRow(query, req.UserKey, req.ResourceCategoryKey, req.ResourceKey).Scan(&response.HasAccess)

		if err != nil {
			return response, HandleDatabaseError(err, r.log, "checking subscripton access")
		}
	} else if req.Type == "tvod" {
		// check for tvod access
		query := `SELECT EXISTS(
			SELECT 1 FROM tvodacces WHERE user_key=$1 AND resource_key=$2
		);`

		err := r.db.Db.QueryRow(query, req.UserKey, req.ResourceKey).Scan(&response.HasAccess)
		if err != nil {
			return response, HandleDatabaseError(err, r.log, "checking subscripton access")
		}
	}

	return response, nil
}

func (r *userSubscription) GetUserSubscriptions(ctx context.Context, req *pb.GetUserSubscriptionsRequest) (*pb.GetUserSubscriptionsResponse, error) {
	var (
		res            = &pb.GetUserSubscriptionsResponse{}
		whereCondition = squirrel.And{}
	)
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 0 {
		req.Limit = 10
	}

	if req.Active != 0 {
		whereCondition = append(whereCondition, squirrel.Eq{"us.active": req.Active > 0})
	}

	if req.Active != 0 {
		whereCondition = append(whereCondition, squirrel.Eq{"sc.visible": req.Visible > 0})
	}

	if req.UserKey != "" {
		whereCondition = append(whereCondition, squirrel.Eq{"us.user_key": req.UserKey})
	}

	if req.FromDate != "" {
		t, err := ParseTimeString(req.FromDate)
		if err != nil {
			whereCondition = append(whereCondition, squirrel.Eq{"us.start_time": t})
		}
	}

	if req.ToDate != "" {
		t, err := ParseTimeString(req.ToDate)
		if err != nil {
			whereCondition = append(whereCondition, squirrel.Eq{"us.end_time": t})
		}
	}

	query := r.db.Builder.Select(`
		us.id, us.user_key, us.start_time, us.end_time, us.active,
		sc.id, sc.title_uz, sc.title_ru, sc.title_en, 
		sc.description_uz, sc.description_ru, sc.description_en,
		sc.image_uz, sc.image_ru, sc.image_en
	`).From("user_subscriptions us").
		Join("subscriptions s ON s.id=us.subscription_id").
		Join("subscription_categories sc ON sc.id=s.category_id").
		Where(whereCondition).OrderBy("us.start_time DESC").
		Offset(uint64((req.Page - 1) * req.Limit)).Limit(uint64(req.Limit))

	rows, err := query.RunWith(r.db.Db).Query()
	if err != nil {
		return res, HandleDatabaseError(err, r.log, "getting user subscriptions")
	}
	defer rows.Close()

	for rows.Next() {
		temp := &pb.UserSubscription{
			Category: &pb.SubscriptionCategory{},
		}

		err := rows.Scan(
			&temp.Id, &temp.UserKey, &temp.StartDate, &temp.EndDate, &temp.Active,
			&temp.Category.Id, &temp.Category.TitleUz, &temp.Category.TitleRu, &temp.Category.TitleEn,
			&temp.Category.DescriptionUz, &temp.Category.DescriptionRu, &temp.Category.DescriptionEn,
			&temp.Category.ImageUz, &temp.Category.ImageRu, &temp.Category.ImageEn,
		)
		if err != nil {
			return res, HandleDatabaseError(err, r.log, "scan user subscriptions")
		}
		res.Items = append(res.Items, temp)
	}

	return res, nil
}
