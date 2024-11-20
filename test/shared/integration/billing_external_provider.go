package integration

import (
	"database/sql"

	"github.com/charmingruby/txgo/internal/billing/core/model"
	"github.com/charmingruby/txgo/internal/billing/database/mysql"
	"github.com/charmingruby/txgo/test/billing/factory"
)

func SubscribeOnPlan(db *sql.DB, email string) error {
	subscriptionRepo := mysql.NewSubscriptionRepository(db)
	planRepo := mysql.NewPlanRepository(db)

	plan, err := factory.MakePlan(planRepo, model.NewPlanFromInput{})
	if err != nil {
		return err
	}

	_, err = factory.MakeSubscription(subscriptionRepo, model.NewSubscriptionFromInput{
		PlanID: plan.ID(),
		Email:  email,
	})

	return err
}
