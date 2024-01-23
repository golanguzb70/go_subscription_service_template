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

type resourceCategory struct {
	db  *db.Postgres
	log logger.Logger
}

func NewCategoryRepo(db *db.Postgres, log logger.Logger) repo.ResourceCategoryI {
	return &resourceCategory{
		db:  db,
		log: log,
	}
}

func (r *resourceCategory) Create(ctx context.Context, req *pb.ResourceCategory) (*pb.ResourceCategory, error) {
	query := r.db.Builder.Insert("resource_categories").
		Columns(`id, title, category_key, allow_all_resources`).
		Values(req.Id, req.Title, req.Key, req.AllowAllResources).
		Suffix("RETURNING created_at, updated_at")

	err := query.RunWith(r.db.Db).QueryRow().Scan(&req.CreatedAt, &req.UpdatedAt)
	if err != nil {
		return req, HandleDatabaseError(err, r.log, "create resource category")
	}

	return req, nil
}

func (r *resourceCategory) Get(ctx context.Context, req *pb.Id) (*pb.ResourceCategory, error) {
	res := &pb.ResourceCategory{}

	query := r.db.Builder.Select("id, title, category_key, allow_all_resources, created_at, updated_at").From("resource_categories")

	if req.Id != "" {
		query = query.Where(squirrel.Eq{"id": req.Id})
	} else {
		return nil, fmt.Errorf("id is required")
	}

	err := query.RunWith(r.db.Db).QueryRow().Scan(
		&res.Id, &res.Title, &res.Key, &res.AllowAllResources,
		&res.CreatedAt, &res.UpdatedAt,
	)

	return res, HandleDatabaseError(err, r.log, "getting resource category")
}

func (r *resourceCategory) Find(ctx context.Context, req *pb.GetListFilter) (*pb.ResourceCategories, error) {
	var (
		res            = &pb.ResourceCategories{}
		whereCondition = PrepareWhere(req.Filters)
		orderBy        = PrepareOrder(req.Sorts)
	)

	query := r.db.Builder.Select("id, title, category_key, allow_all_resources, created_at, updated_at").
		From("resource_categories").
		Where(whereCondition).
		OrderBy(orderBy).Offset(uint64((req.Page - 1) * req.Limit)).Limit(uint64(req.Limit))

	rows, err := query.RunWith(r.db.Db).Query()
	if err != nil {
		return nil, HandleDatabaseError(err, r.log, "error while finding resource categories")
	}
	defer rows.Close()

	for rows.Next() {
		temp := &pb.ResourceCategory{}
		err = rows.Scan(
			&temp.Id, &temp.Title, &temp.Key, &temp.AllowAllResources,
			&temp.CreatedAt, &temp.UpdatedAt,
		)
		if err != nil {
			return nil, HandleDatabaseError(err, r.log, "error while scanning resource_category")
		}

		res.Items = append(res.Items, temp)
	}

	count := r.db.Builder.Select("count(1)").
		From("resource_categories").
		Where(whereCondition)

	err = count.RunWith(r.db.Db).Scan(&res.Count)

	return res, HandleDatabaseError(err, r.log, "getting resource category count")
}

func (r *resourceCategory) Update(ctx context.Context, req *pb.ResourceCategory) (*pb.ResourceCategory, error) {
	var (
		mp             = make(map[string]interface{})
		whereCondition = squirrel.And{squirrel.Eq{"id": req.Id}}
	)

	mp["title"] = req.Title
	mp["category_key"] = req.Key
	mp["allow_all_resources"] = req.AllowAllResources
	mp["updated_at"] = time.Now()

	query := r.db.Builder.Update("resource_categories").SetMap(mp).
		Where(whereCondition).
		Suffix("RETURNING updated_at, created_at")

	err := query.RunWith(r.db.Db).QueryRow().Scan(
		&req.CreatedAt, &req.UpdatedAt,
	)

	return req, HandleDatabaseError(err, r.log, "update resource_category")
}

func (r *resourceCategory) Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	query := r.db.Builder.Delete("resource_categories").Where(squirrel.Eq{"id": req.Id})
	_, err := query.RunWith(r.db.Db).Exec()
	return &pb.Empty{}, HandleDatabaseError(err, r.log, "delete from resource categories")
}
