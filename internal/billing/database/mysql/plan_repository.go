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
	PLANS_TABLE = "plans"
)

func NewPlanRepository(db mysql.Database) *PlanRepository {
	return &PlanRepository{db: db}
}

type PlanRepository struct {
	db mysql.Database
}

type planRow struct {
	ID              string
	Name            string
	Description     string
	Amount          int
	Periodicity     string
	TrialPeriodDays int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (r *PlanRepository) FindByName(name string) (*model.Plan, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = ?", PLANS_TABLE)

	queryResult := r.db.QueryRow(query, name)

	var row planRow

	if err := queryResult.Scan(
		&row.ID,
		&row.Name,
		&row.Description,
		&row.Amount,
		&row.Periodicity,
		&row.TrialPeriodDays,
		&row.CreatedAt,
		&row.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, core_err.NewPersistenceErr(err, "plan find_by_name", "mysql")
	}

	return r.mapToDomain(row), nil
}

func (r *PlanRepository) FindByID(id string) (*model.Plan, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", PLANS_TABLE)

	queryResult := r.db.QueryRow(query, id)

	var row planRow

	if err := queryResult.Scan(
		&row.ID,
		&row.Name,
		&row.Description,
		&row.Amount,
		&row.Periodicity,
		&row.TrialPeriodDays,
		&row.CreatedAt,
		&row.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, core_err.NewPersistenceErr(err, "plan find_by_id", "mysql")
	}

	return r.mapToDomain(row), nil
}

func (r *PlanRepository) Store(plan *model.Plan) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name, description, amount, periodicity, trial_period_days, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", PLANS_TABLE)

	_, err := r.db.Exec(query,
		plan.ID(),
		plan.Name(),
		plan.Description(),
		plan.Amount(),
		plan.Periodicity(),
		plan.TrialPeriodDays(),
		plan.CreatedAt(),
		plan.UpdatedAt())
	if err != nil {
		return core_err.NewPersistenceErr(err, "plan store", "mysql")
	}

	return nil
}

func (r *PlanRepository) mapToDomain(plan planRow) *model.Plan {
	return model.NewPlanFrom(model.NewPlanFromInput{
		ID:              plan.ID,
		Name:            plan.Name,
		Amount:          plan.Amount,
		Periodicity:     plan.Periodicity,
		Description:     plan.Description,
		TrialPeriodDays: plan.TrialPeriodDays,
		CreatedAt:       plan.CreatedAt,
		UpdatedAt:       plan.UpdatedAt,
	})
}
