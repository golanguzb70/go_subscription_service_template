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

type subscriptionCategory struct {
	db  *db.Postgres
	log logger.Logger
}

func NewSubscriptionCategoryRepo(db *db.Postgres, log logger.Logger) repo.SubscriptionCategoryI {
	return &subscriptionCategory{
		db:  db,
		log: log,
	}
}

func (r *subscriptionCategory) Create(ctx context.Context, req *pb.SubscriptionCategory) (*pb.SubscriptionCategory, error) {
	fmt.Println(req)
	query := r.db.Builder.Insert("subscription_categories").
		Columns(`
			id, title_uz, title_ru, title_en, description_uz, description_ru, description_en,
			image_uz, image_ru, image_en, active, visible
			`).
		Values(req.Id, req.TitleUz, req.TitleRu, req.TitleEn, req.DescriptionUz, req.DescriptionRu, req.DescriptionEn,
			req.ImageUz, req.ImageRu, req.ImageEn, req.Active, req.Visible).
		Suffix("RETURNING created_at, updated_at")

	err := query.RunWith(r.db.Db).QueryRow().Scan(&req.CreatedAt, &req.UpdatedAt)
	if err != nil {
		return req, HandleDatabaseError(err, r.log, "create subscription category")
	}

	return req, nil
}

func (r *subscriptionCategory) Get(ctx context.Context, req *pb.Id) (*pb.SubscriptionCategory, error) {
	res := &pb.SubscriptionCategory{}

	query := r.db.Builder.Select(`
		id, title_uz, title_ru, title_en, description_uz, description_ru, description_en,
		image_uz, image_ru, image_en, active, visible, created_at, updated_at
	`).From("subscription_categories")

	if req.Id != "" {
		query = query.Where(squirrel.Eq{"id": req.Id})
	} else {
		return nil, fmt.Errorf("id is required")
	}

	err := query.RunWith(r.db.Db).QueryRow().Scan(
		&res.Id, &res.TitleUz, &res.TitleRu, &res.TitleEn, &res.DescriptionUz, &res.DescriptionRu, &res.DescriptionEn,
		&res.ImageUz, &res.ImageRu, &res.ImageEn, &res.Active, &res.Visible, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.log, "getting subscription category")
	}

	query = r.db.Builder.Select("r.id, r.title, r.category_key, r.allow_all_resources, r.created_at, r.updated_at").
		From("resource_categories r").
		InnerJoin("resource_subsription_categories rs ON rs.resource_category_id = r.id").
		Where(squirrel.Eq{"subscription_category_id": res.Id}).
		OrderBy("r.title ASC")

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

		res.ResourceCategories = append(res.ResourceCategories, temp)
	}

	return res, HandleDatabaseError(err, r.log, "getting subscription category")
}

func (r *subscriptionCategory) Find(ctx context.Context, req *pb.GetListFilter) (*pb.SubscriptionCategories, error) {
	var (
		res            = &pb.SubscriptionCategories{}
		whereCondition = PrepareWhere(req.Filters)
		orderBy        = PrepareOrder(req.Sorts)
	)

	query := r.db.Builder.Select(`
		id, title_uz, title_ru, title_en, description_uz, description_ru, description_en,
		image_uz, image_ru, image_en, active, visible, created_at, updated_at
	`).From("subscription_categories")

	if len(req.Filters) > 0 {
		query = query.Where(whereCondition)

	}
	if len(req.Sorts) > 0 {
		query = query.OrderBy(orderBy)
	}

	query = query.Offset(uint64((req.Page - 1) * req.Limit)).Limit(uint64(req.Limit))

	rows, err := query.RunWith(r.db.Db).Query()
	if err != nil {
		return nil, HandleDatabaseError(err, r.log, "error while finding subscription categories")
	}
	defer rows.Close()

	for rows.Next() {
		temp := &pb.SubscriptionCategory{}
		err = rows.Scan(
			&temp.Id, &temp.TitleUz, &temp.TitleRu, &temp.TitleEn, &temp.DescriptionUz, &temp.DescriptionRu, &temp.DescriptionEn,
			&temp.ImageUz, &temp.ImageRu, &temp.ImageEn, &temp.Active, &temp.Visible, &temp.CreatedAt, &temp.UpdatedAt,
		)
		if err != nil {
			return nil, HandleDatabaseError(err, r.log, "error while scanning resource_category")
		}

		res.Items = append(res.Items, temp)
	}

	count := r.db.Builder.Select("count(1)").
		From("subscription_categories").
		Where(whereCondition)

	err = count.RunWith(r.db.Db).Scan(&res.Count)

	return res, HandleDatabaseError(err, r.log, "getting resource category count")
}

func (r *subscriptionCategory) Update(ctx context.Context, req *pb.SubscriptionCategory) (*pb.SubscriptionCategory, error) {
	var (
		mp             = make(map[string]interface{})
		whereCondition = squirrel.And{squirrel.Eq{"id": req.Id}}
	)

	mp["id"] = req.Id
	mp["title_uz"] = req.TitleUz
	mp["title_ru"] = req.TitleRu
	mp["title_en"] = req.TitleEn
	mp["description_uz"] = req.DescriptionUz
	mp["description_ru"] = req.DescriptionRu
	mp["description_en"] = req.DescriptionEn
	mp["image_uz"] = req.ImageUz
	mp["image_ru"] = req.ImageRu
	mp["image_en"] = req.ImageEn
	mp["active"] = req.Active
	mp["visible"] = req.Visible

	mp["updated_at"] = time.Now()

	query := r.db.Builder.Update("subscription_categories").SetMap(mp).
		Where(whereCondition).
		Suffix("RETURNING updated_at, created_at")

	err := query.RunWith(r.db.Db).QueryRow().Scan(
		&req.CreatedAt, &req.UpdatedAt,
	)

	return req, HandleDatabaseError(err, r.log, "update subscription category")
}

func (r *subscriptionCategory) Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	query := r.db.Builder.Delete("subscription_categories").Where(squirrel.Eq{"id": req.Id})
	_, err := query.RunWith(r.db.Db).Exec()
	return &pb.Empty{}, HandleDatabaseError(err, r.log, "delete from resource categories")
}

func (r *subscriptionCategory) AddResourceCategory(ctx context.Context, req *pb.SubscriptionResourceCategoryIds) (*pb.Empty, error) {
	query := r.db.Builder.Insert("resource_subsription_categories").
		Columns("id, subscription_category_id, resource_category_id")

	for _, e := range req.ResourceCategoryId {
		query = query.Values(uuid.New().String(), req.SubscriptionCategoryId, e)
	}

	_, err := query.RunWith(r.db.Db).Exec()

	return &pb.Empty{}, HandleDatabaseError(err, r.log, "add resource category to subscription category")
}

func (r *subscriptionCategory) RemoveResourceCategory(ctx context.Context, req *pb.SubscriptionResourceCategoryIds) (*pb.Empty, error) {
	query := r.db.Builder.Delete("resource_subsription_categories").
		Where(squirrel.Eq{"subscription_category_id": req.SubscriptionCategoryId}).Where(squirrel.Eq{"resource_category_id": req.ResourceCategoryId})

	_, err := query.RunWith(r.db.Db).Exec()

	return &pb.Empty{}, HandleDatabaseError(err, r.log, "remove resource category from subscription category")
}
