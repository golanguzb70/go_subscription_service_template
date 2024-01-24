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
	"github.com/google/uuid"
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

	if err != nil {
		return res, HandleDatabaseError(err, r.log, "getting resource category")
	}

	query = r.db.Builder.Select("r.id, r.title, r.resource_key, r.created_at, r.updated_at").From("resources_categories_m2m").
		InnerJoin("resources as r ON r.id=resource_id").Where(squirrel.Eq{"category_id": res.Id})

	rows, err := query.RunWith(r.db.Db).Query()
	if err != nil {
		return nil, HandleDatabaseError(err, r.log, "error while getting resources of category")
	}
	defer rows.Close()

	for rows.Next() {
		temp := &pb.Resource{}

		err = rows.Scan(
			&temp.Id, &temp.Title, &temp.Key, &temp.CreatedAt, &temp.UpdatedAt,
		)
		if err != nil {
			return nil, HandleDatabaseError(err, r.log, "error while scanning resource")
		}

		res.Resources = append(res.Resources, temp)
	}

	return res, HandleDatabaseError(err, r.log, "getting resource category")
}

func (r *resourceCategory) Find(ctx context.Context, req *pb.GetListFilter) (*pb.ResourceCategories, error) {
	var (
		res            = &pb.ResourceCategories{}
		whereCondition = PrepareWhere(req.Filters)
		orderBy        = PrepareOrder(req.Sorts)
	)

	query := r.db.Builder.Select("id, title, category_key, allow_all_resources, created_at, updated_at").
		From("resource_categories")
	if len(req.Filters) > 0 {
		query = query.Where(whereCondition)

	}
	if len(req.Sorts) > 0 {
		query = query.OrderBy(orderBy)
	}

	query = query.Offset(uint64((req.Page - 1) * req.Limit)).Limit(uint64(req.Limit))

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

func (r *resourceCategory) AddResource(ctx context.Context, req *pb.ResourceAndCategoryIds) (*pb.Empty, error) {
	query := r.db.Builder.Insert("resources_categories_m2m").
		Columns("id, category_id, resource_id")

	for _, e := range req.ResourceId {
		query = query.Values(uuid.New().String(), req.CategoryId, e)
	}

	_, err := query.RunWith(r.db.Db).Exec()

	return &pb.Empty{}, HandleDatabaseError(err, r.log, "add resource to categories")
}

func (r *resourceCategory) RemoveResource(ctx context.Context, req *pb.ResourceAndCategoryIds) (*pb.Empty, error) {
	query := r.db.Builder.Delete("resources_categories_m2m").
		Where(squirrel.Eq{"category_id": req.CategoryId}).Where(squirrel.Eq{"resource_id": req.ResourceId})

	_, err := query.RunWith(r.db.Db).Exec()

	return &pb.Empty{}, HandleDatabaseError(err, r.log, "delete from resource categories")
}
