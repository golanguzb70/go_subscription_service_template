package postgres

import (
	"github.com/golanguzb70/go_subscription_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type resourceCategory struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) repo.ResourceCategoryI {
	return &resourceCategory{
		db: db,
	}
}
