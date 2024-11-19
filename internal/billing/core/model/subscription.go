package model

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
)

const (
	SUBSCRIPTION_STATUS_CANCELED = "CANCELED"
	SUBSCRIPTION_STATUS_EXPIRED  = "EXPIRED"
	SUBSCRIPTION_STATUS_PENDING  = "PENDING"
	SUBSCRIPTION_STATUS_ACTIVE   = "ACTIVE"
)

type Subscription struct {
	id        string
	email     string
	planID    string
	status    string
	startDate time.Time
	endDate   *time.Time
	autoRenew bool
	createdAt time.Time
	updateAt  time.Time
}

type NewSubscriptionInput struct {
	Email     string
	PlanID    string
	StartDate time.Time
	EndDate   *time.Time
	AutoRenew bool
}

func NewSubscription(input NewSubscriptionInput) *Subscription {
	return &Subscription{
		id:        core.NewID(),
		email:     input.Email,
		planID:    input.PlanID,
		status:    SUBSCRIPTION_STATUS_PENDING,
		startDate: input.StartDate,
		endDate:   input.EndDate,
		autoRenew: input.AutoRenew,
		createdAt: time.Now(),
		updateAt:  time.Now(),
	}
}
