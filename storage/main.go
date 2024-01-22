package storage

import (
	"github.com/golanguzb70/go_subscription_service/storage/postgres"
	"github.com/golanguzb70/go_subscription_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	ResourceCategory() repo.ResourceCategoryI
}

type storagePg struct {
	resourceCategory repo.ResourceCategoryI
}

func New(db *sqlx.DB) StorageI {
	return &storagePg{
		resourceCategory: postgres.NewCategoryRepo(db),
	}
}

func (s *storagePg) ResourceCategory() repo.ResourceCategoryI {
	return s.resourceCategory
}
