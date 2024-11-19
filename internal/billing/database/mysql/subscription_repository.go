package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/charmingruby/txgo/internal/billing/core/model"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/charmingruby/txgo/internal/shared/database/mysql"
)

const (
	SUBSCRIPTIONS_TABLE = "subscriptions"
)

func NewSubscriptionRepository(db mysql.Database) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

type SubscriptionRepository struct {
	db mysql.Database
}

type subscriptionRow struct {
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

func (r *SubscriptionRepository) FindActiveByEmail(email string) (*model.Subscription, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = ? AND status = ?", SUBSCRIPTIONS_TABLE)

	queryResult := r.db.QueryRow(query, email, model.SUBSCRIPTION_STATUS_ACTIVE)

	var row subscriptionRow

	if err := queryResult.Scan(
		&row.ID,
		&row.Email,
		&row.PlanID,
		&row.Status,
		&row.StartDate,
		&row.EndDate,
		&row.AutoRenew,
		&row.CreatedAt,
		&row.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, core_err.NewPersistenceErr(err, "subscription find_active_by_email", "mysql")
	}

	return r.mapToDomain(row), nil
}

func (r *SubscriptionRepository) FindNonInactiveByEmailAndPlanID(email, planID string) (*model.Subscription, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ? AND plan_id = ? AND status IN (?, ?)", SUBSCRIPTIONS_TABLE)

	queryResult := r.db.QueryRow(query, email, planID, model.SUBSCRIPTION_STATUS_PENDING, model.SUBSCRIPTION_STATUS_ACTIVE)

	var row subscriptionRow

	if err := queryResult.Scan(
		&row.ID,
		&row.Email,
		&row.PlanID,
		&row.Status,
		&row.StartDate,
		&row.EndDate,
		&row.AutoRenew,
		&row.CreatedAt,
		&row.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, core_err.NewPersistenceErr(err, "subscription find_by_id", "mysql")
	}

	return r.mapToDomain(row), nil
}

func (r *SubscriptionRepository) Store(subscription *model.Subscription) error {
	query := fmt.Sprintf("INSERT INTO %s (id, email, plan_id, status, start_date, end_date, auto_renew, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", SUBSCRIPTIONS_TABLE)

	_, err := r.db.Exec(query,
		subscription.ID(),
		subscription.Email(),
		subscription.PlanID(),
		subscription.Status(),
		subscription.StartDate(),
		subscription.EndDate(),
		subscription.AutoRenew(),
		subscription.CreatedAt(),
		subscription.UpdatedAt())
	if err != nil {
		return core_err.NewPersistenceErr(err, "subscription store", "mysql")
	}

	return nil
}

func (r *SubscriptionRepository) mapToDomain(subscription subscriptionRow) *model.Subscription {
	return model.NewSubscriptionFrom(model.NewSubscriptionFromInput{
		ID:        subscription.ID,
		Email:     subscription.Email,
		PlanID:    subscription.PlanID,
		Status:    subscription.Status,
		StartDate: subscription.StartDate,
		EndDate:   subscription.EndDate,
		AutoRenew: subscription.AutoRenew,
		CreatedAt: subscription.CreatedAt,
		UpdatedAt: subscription.UpdatedAt,
	})
}
