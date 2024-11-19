package service

import (
	"time"

	"github.com/charmingruby/txgo/internal/billing/core/model"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type SubscribeOnPlanInput struct {
	Email     string
	PlanID    string
	AutoRenew bool
}

func (s *Service) SubscribeOnPlanService(params SubscribeOnPlanInput) error {
	plan, err := s.planRepo.FindByID(params.PlanID)
	if err != nil {
		return err
	}

	if plan == nil {
		return core_err.NewResourceNotFoundErr("plan")
	}

	subscriptionFound, err := s.subscriptionRepo.FindNonInactiveByEmailAndPlanID(params.Email, params.PlanID)
	if err != nil {
		return err
	}

	if subscriptionFound != nil {
		return core_err.NewResourceAlreadyExistsErr("subscription")
	}

	subscription := model.NewSubscription(model.NewSubscriptionInput{
		Email:     params.Email,
		PlanID:    params.PlanID,
		StartDate: time.Now(),
		EndDate:   calculateEndDate(time.Now(), plan.Periodicity()),
		AutoRenew: params.AutoRenew,
	})

	if err := s.subscriptionRepo.Store(subscription); err != nil {
		return err
	}

	return nil
}

func calculateEndDate(startDate time.Time, periodicity string) *time.Time {
	if periodicity == model.MONTHLY_PLAN_PERIODICITY {
		endDate := startDate.AddDate(0, 1, 0)
		return &endDate
	}

	endDate := startDate.AddDate(1, 0, 0)

	return &endDate
}
