package repository

import "github.com/charmingruby/txgo/internal/billing/core/model"

type SubscriptionRepository interface {
	Store(subscription *model.Subscription) error
	FindActiveByEmail(email string) (*model.Subscription, error)
	FindNonInactiveByEmailAndPlanID(email, planID string) (*model.Subscription, error)
}
