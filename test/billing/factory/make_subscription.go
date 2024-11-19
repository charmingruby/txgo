package factory

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/charmingruby/txgo/internal/billing/core/model"
	"github.com/charmingruby/txgo/internal/billing/core/repository"
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/helper"
)

func MakeSubscription(subscriptionRepo repository.SubscriptionRepository, params model.NewSubscriptionFromInput) (*model.Subscription, error) {
	subscription := createSubscription(params)

	if err := subscriptionRepo.Store(subscription); err != nil {
		return nil, err
	}

	return subscription, nil
}

func createSubscription(params model.NewSubscriptionFromInput) *model.Subscription {
	now := time.Now()

	input := model.NewSubscriptionFromInput{
		ID:        helper.If[string](params.ID != "", params.ID, core.NewID()),
		Email:     helper.If[string](params.Email != "", params.Email, gofakeit.Email()),
		PlanID:    params.PlanID,
		Status:    helper.If[string](params.Status != "", params.Status, model.SUBSCRIPTION_STATUS_ACTIVE),
		StartDate: helper.If[time.Time](params.StartDate != time.Time{}, params.StartDate, now),
		EndDate:   helper.If[*time.Time](params.EndDate != nil, params.EndDate, nil),
		AutoRenew: params.AutoRenew,
		CreatedAt: helper.If[time.Time](params.CreatedAt != time.Time{}, params.CreatedAt, now),
		UpdatedAt: helper.If[time.Time](params.UpdatedAt != time.Time{}, params.UpdatedAt, now),
	}

	return model.NewSubscriptionFrom(input)
}
