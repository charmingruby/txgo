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

type NewSubscriptionFromInput struct {
	ID        string
	Email     string
	PlanID    string
	Status    string
	StartDate time.Time
	EndDate   *time.Time
	AutoRenew bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSubscriptionFrom(input NewSubscriptionFromInput) *Subscription {
	return &Subscription{
		id:        input.ID,
		email:     input.Email,
		planID:    input.PlanID,
		status:    input.Status,
		startDate: input.StartDate,
		endDate:   input.EndDate,
		autoRenew: input.AutoRenew,
		createdAt: input.CreatedAt,
		updateAt:  input.UpdatedAt,
	}
}

func (s *Subscription) ID() string {
	return s.id
}

func (s *Subscription) Email() string {
	return s.email
}

func (s *Subscription) PlanID() string {
	return s.planID
}

func (s *Subscription) Status() string {
	return s.status
}

func (s *Subscription) StartDate() time.Time {
	return s.startDate
}

func (s *Subscription) EndDate() *time.Time {
	return s.endDate
}

func (s *Subscription) AutoRenew() bool {
	return s.autoRenew
}

func (s *Subscription) CreatedAt() time.Time {
	return s.createdAt
}

func (s *Subscription) UpdatedAt() time.Time {
	return s.updateAt
}
