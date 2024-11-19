package repository

import "github.com/charmingruby/txgo/internal/billing/core/model"

type SubscriptionRepository interface {
	Store(subscription *model.Subscription) error
	FindByEmail(email string) (*model.Subscription, error)
}
