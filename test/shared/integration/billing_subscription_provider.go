package integration

import (
	"database/sql"

	"github.com/charmingruby/txgo/internal/billing/database/mysql"
	"github.com/charmingruby/txgo/internal/billing/integration/provider"
	"github.com/charmingruby/txgo/internal/shared/integration"
)

func NewBillingSubscriptionProvider(db *sql.DB) integration.BillingSubscriptionStatusIntegration {
	subscriptionRepo := mysql.NewSubscriptionRepository(db)
	return provider.NewPublic(subscriptionRepo)
}
