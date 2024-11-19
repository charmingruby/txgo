package repository

import "github.com/charmingruby/txgo/internal/billing/core/model"

type PlanRepository interface {
	Store(plan *model.Plan) error
	FindByID(id string) (*model.Subscription, error)
}
