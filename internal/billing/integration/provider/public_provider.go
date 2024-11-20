package provider

import (
	"fmt"
	"log/slog"

	"github.com/charmingruby/txgo/internal/billing/core/repository"
)

func NewPublic(subscriptionRepo repository.SubscriptionRepository) *Public {
	return &Public{
		subscriptionRepo: subscriptionRepo,
	}
}

type Public struct {
	subscriptionRepo repository.SubscriptionRepository
}

func (a *Public) IsSubscriptionActive(email string) bool {
	subscription, err := a.subscriptionRepo.FindActiveByEmail(email)
	if err != nil {
		slog.Info(fmt.Sprintf("[BILLING PUBLIC PROVIDER] %s", err.Error()))
		return false
	}

	return subscription != nil
}
