package storage

import (
	"github.com/golanguzb70/go_subscription_service/pkg/db"
	"github.com/golanguzb70/go_subscription_service/pkg/logger"
	"github.com/golanguzb70/go_subscription_service/storage/postgres"
	"github.com/golanguzb70/go_subscription_service/storage/repo"
)

type StorageI interface {
	ResourceCategory() repo.ResourceCategoryI
	SubscriptionCategory() repo.SubscriptionCategoryI
	Resource() repo.ResourceI
	Subscription() repo.SubscriptionI
	UserSubscription() repo.UserSubscriptionI
}

type storagePg struct {
	resourceCategory     repo.ResourceCategoryI
	subscriptionCategory repo.SubscriptionCategoryI
	resource             repo.ResourceI
	subscription         repo.SubscriptionI
	userSubscription     repo.UserSubscriptionI
}

func New(db *db.Postgres, log logger.Logger) StorageI {
	return &storagePg{
		resourceCategory:     postgres.NewCategoryRepo(db, log),
		subscriptionCategory: postgres.NewSubscriptionCategoryRepo(db, log),
		resource:             postgres.NewResourceRepo(db, log),
		subscription:         postgres.NewSubscriptionRepo(db, log),
		userSubscription:     postgres.NewUserSubscriptionRepo(db, log),
	}
}

func (s *storagePg) ResourceCategory() repo.ResourceCategoryI {
	return s.resourceCategory
}

func (s *storagePg) Resource() repo.ResourceI {
	return s.resource
}

func (s *storagePg) SubscriptionCategory() repo.SubscriptionCategoryI {
	return s.subscriptionCategory
}

func (s *storagePg) Subscription() repo.SubscriptionI {
	return s.subscription
}

func (s *storagePg) UserSubscription() repo.UserSubscriptionI {
	return s.userSubscription
}