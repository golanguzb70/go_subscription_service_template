package storage

import (
	"github.com/golanguzb70/go_subscription_service/pkg/db"
	"github.com/golanguzb70/go_subscription_service/pkg/logger"
	"github.com/golanguzb70/go_subscription_service/storage/postgres"
	"github.com/golanguzb70/go_subscription_service/storage/repo"
)

type StorageI interface {
	ResourceCategory() repo.ResourceCategoryI
}

type storagePg struct {
	resourceCategory repo.ResourceCategoryI
}

func New(db *db.Postgres, log logger.Logger) StorageI {
	return &storagePg{
		resourceCategory: postgres.NewCategoryRepo(db, log),
	}
}

func (s *storagePg) ResourceCategory() repo.ResourceCategoryI {
	return s.resourceCategory
}
