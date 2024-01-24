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

type resource struct {
	db  *db.Postgres
	log logger.Logger
}

func NewResourceRepo(db *db.Postgres, log logger.Logger) repo.ResourceI {
	return &resource{
		db:  db,
		log: log,
	}
}

func (r *resource) Create(ctx context.Context, req *pb.Resource) (*pb.Resource, error) {
	query := r.db.Builder.Insert("resources").
		Columns(`id, title, resource_key`).
		Values(req.Id, req.Title, req.Key).
		Suffix("RETURNING created_at, updated_at")

	err := query.RunWith(r.db.Db).QueryRow().Scan(&req.CreatedAt, &req.UpdatedAt)
	if err != nil {
		return req, HandleDatabaseError(err, r.log, "create resource")
	}

	return req, nil
}

func (r *resource) Get(ctx context.Context, req *pb.Id) (*pb.Resource, error) {
	res := &pb.Resource{}

	query := r.db.Builder.Select("id, title, resource_key, created_at, updated_at").From("resources")

	if req.Id != "" {
		query = query.Where(squirrel.Eq{"id": req.Id})
	} else {
		return nil, fmt.Errorf("id is required")
	}

	err := query.RunWith(r.db.Db).QueryRow().Scan(
		&res.Id, &res.Title, &res.Key,
		&res.CreatedAt, &res.UpdatedAt,
	)

	return res, HandleDatabaseError(err, r.log, "getting resource")
}

func (r *resource) Find(ctx context.Context, req *pb.GetListFilter) (*pb.Resources, error) {
	var (
		res            = &pb.Resources{}
		whereCondition = PrepareWhere(req.Filters)
		orderBy        = PrepareOrder(req.Sorts)
	)

	query := r.db.Builder.Select("id, title, resource_key, created_at, updated_at").
		From("resources")
		
	if len(req.Filters) > 0 {
		query = query.Where(whereCondition)

	}
	if len(req.Sorts) > 0 {
		query = query.OrderBy(orderBy)
	}

	query = query.Offset(uint64((req.Page - 1) * req.Limit)).Limit(uint64(req.Limit))

	rows, err := query.RunWith(r.db.Db).Query()
	if err != nil {
		return nil, HandleDatabaseError(err, r.log, "error while finding resource")
	}
	defer rows.Close()

	for rows.Next() {
		temp := &pb.Resource{}
		err = rows.Scan(
			&temp.Id, &temp.Title, &temp.Key,
			&temp.CreatedAt, &temp.UpdatedAt,
		)
		if err != nil {
			return nil, HandleDatabaseError(err, r.log, "error while scanning resource")
		}

		res.Items = append(res.Items, temp)
	}

	count := r.db.Builder.Select("count(1)").
		From("resources").
		Where(whereCondition)

	err = count.RunWith(r.db.Db).Scan(&res.Count)

	return res, HandleDatabaseError(err, r.log, "getting resource count")
}

func (r *resource) Update(ctx context.Context, req *pb.Resource) (*pb.Resource, error) {
	var (
		mp             = make(map[string]interface{})
		whereCondition = squirrel.And{squirrel.Eq{"id": req.Id}}
	)

	mp["title"] = req.Title
	mp["resource_key"] = req.Key
	mp["updated_at"] = time.Now()

	query := r.db.Builder.Update("resources").SetMap(mp).
		Where(whereCondition).
		Suffix("RETURNING updated_at, created_at")

	err := query.RunWith(r.db.Db).QueryRow().Scan(
		&req.CreatedAt, &req.UpdatedAt,
	)

	return req, HandleDatabaseError(err, r.log, "update resource")
}

func (r *resource) Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	query := r.db.Builder.Delete("resources").Where(squirrel.Eq{"id": req.Id})
	_, err := query.RunWith(r.db.Db).Exec()
	return &pb.Empty{}, HandleDatabaseError(err, r.log, "delete from resource")
}
