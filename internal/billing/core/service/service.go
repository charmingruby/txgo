package service

import "github.com/charmingruby/txgo/internal/billing/core/repository"

type Service struct {
	subscriptionRepo repository.SubscriptionRepository
	planRepo         repository.PlanRepository
}

func New(
	subscriptionRepo repository.SubscriptionRepository,
	planRepo repository.PlanRepository,
) *Service {
	return &Service{
		subscriptionRepo: subscriptionRepo,
		planRepo:         planRepo,
	}
}
