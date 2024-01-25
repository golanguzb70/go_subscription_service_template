package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	pb "github.com/golanguzb70/go_subscription_service/genproto/subscription_service"
	"github.com/golanguzb70/go_subscription_service/pkg/db"
	"github.com/golanguzb70/go_subscription_service/pkg/logger"
	"github.com/golanguzb70/go_subscription_service/storage/repo"
)

type subscription struct {
	db  *db.Postgres
	log logger.Logger
}

func NewSubscriptionRepo(db *db.Postgres, log logger.Logger) repo.SubscriptionI {
	return &subscription{
		db:  db,
		log: log,
	}
}

func (r *subscription) Create(ctx context.Context, req *pb.Subscription) (*pb.Subscription, error) {
	query := r.db.Builder.Insert("subscriptions").
		Columns(`
			id, title_uz, title_ru, title_en, active, 
			price, duration_type, duration, category_id
		`).
		Values(
			req.Id, req.TitleUz, req.TitleRu, req.TitleEn, req.Active,
			req.Price, req.DurationType, req.Duration, req.CategoryId,
		).
		Suffix("RETURNING created_at, updated_at")

	err := query.RunWith(r.db.Db).QueryRow().Scan(&req.CreatedAt, &req.UpdatedAt)
	if err != nil {
		return req, HandleDatabaseError(err, r.log, "create subscription")
	}

	return req, nil
}

func (r *subscription) Get(ctx context.Context, req *pb.Id) (*pb.Subscription, error) {
	res := &pb.Subscription{}

	query := r.db.Builder.Select(`
		id, title_uz, title_ru, title_en, active, 
		price, duration_type, duration, category_id, created_at, updated_at
	`).From("subscriptions")

	if req.Id != "" {
		query = query.Where(squirrel.Eq{"id": req.Id})
	} else {
		return nil, fmt.Errorf("id is required")
	}

	err := query.RunWith(r.db.Db).QueryRow().Scan(
		&res.Id, &res.TitleUz, &res.TitleRu, &res.TitleEn, &res.Active,
		&res.Price, &res.DurationType, &res.Duration, &res.CategoryId,
		&res.CreatedAt, &res.UpdatedAt,
	)

	return res, HandleDatabaseError(err, r.log, "getting subscription")
}

func (r *subscription) Find(ctx context.Context, req *pb.GetListFilter) (*pb.Subscriptions, error) {
	var (
		res            = &pb.Subscriptions{}
		whereCondition = PrepareWhere(req.Filters)
		orderBy        = PrepareOrder(req.Sorts)
	)

	query := r.db.Builder.Select(`
		id, title_uz, title_ru, title_en, active, 
		price, duration_type, duration, category_id, created_at, updated_at
	`).
		From("subscriptions")

	if len(req.Filters) > 0 {
		query = query.Where(whereCondition)

	}
	if len(req.Sorts) > 0 {
		query = query.OrderBy(orderBy)
	}

	query = query.Offset(uint64((req.Page - 1) * req.Limit)).Limit(uint64(req.Limit))

	rows, err := query.RunWith(r.db.Db).Query()
	if err != nil {
		return nil, HandleDatabaseError(err, r.log, "error while finding subscription")
	}
	defer rows.Close()

	for rows.Next() {
		temp := &pb.Subscription{}
		err = rows.Scan(
			&temp.Id, &temp.TitleUz, &temp.TitleRu, &temp.TitleEn, &temp.Active,
			&temp.Price, &temp.DurationType, &temp.Duration, &temp.CategoryId,
			&temp.CreatedAt, &temp.UpdatedAt,
		)
		if err != nil {
			return nil, HandleDatabaseError(err, r.log, "error while scanning subscription")
		}

		res.Items = append(res.Items, temp)
	}

	count := r.db.Builder.Select("count(1)").
		From("subscriptions").
		Where(whereCondition)

	err = count.RunWith(r.db.Db).Scan(&res.Count)

	return res, HandleDatabaseError(err, r.log, "getting subscription count")
}

func (r *subscription) Update(ctx context.Context, req *pb.Subscription) (*pb.Subscription, error) {
	var (
		mp             = make(map[string]interface{})
		whereCondition = squirrel.And{squirrel.Eq{"id": req.Id}}
	)
	mp["title_uz"] = req.TitleUz
	mp["title_ru"] = req.TitleRu
	mp["title_en"] = req.TitleEn
	mp["active"] = req.Active
	mp["price"] = req.Price
	mp["duration_type"] = req.DurationType
	mp["duration"] = req.Duration
	mp["category_id"] = req.CategoryId
	mp["updated_at"] = time.Now()

	query := r.db.Builder.Update("subscriptions").SetMap(mp).
		Where(whereCondition).
		Suffix("RETURNING updated_at, created_at")

	err := query.RunWith(r.db.Db).QueryRow().Scan(
		&req.CreatedAt, &req.UpdatedAt,
	)

	return req, HandleDatabaseError(err, r.log, "update subscription")
}

func (r *subscription) Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	query := r.db.Builder.Delete("subscriptions").Where(squirrel.Eq{"id": req.Id})
	_, err := query.RunWith(r.db.Db).Exec()
	return &pb.Empty{}, HandleDatabaseError(err, r.log, "delete from subscription")
}
