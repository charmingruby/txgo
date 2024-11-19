package model

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

const (
	UNDEFINED_PLAN_PERIODICITY = "undefined"
	MONTHLY_PLAN_PERIODICITY   = "month"
	YEARLY_PLAN_PERIODICITY    = "year"
)

type Plan struct {
	id              string
	name            string
	description     string
	amount          int
	periodicity     string
	trialPeriodDays int
	createdAt       time.Time
	updatedAt       time.Time
}

type NewPlanInput struct {
	Name            string
	Description     string
	Amount          int
	Periodicity     string
	TrialPeriodDays int
}

func NewPlan(input NewPlanInput) *Plan {
	return &Plan{
		id:              core.NewID(),
		name:            input.Name,
		description:     input.Description,
		amount:          input.Amount,
		periodicity:     input.Periodicity,
		trialPeriodDays: input.TrialPeriodDays,
		createdAt:       time.Now(),
		updatedAt:       time.Now(),
	}
}

type NewPlanFromInput struct {
	ID              string
	Name            string
	Description     string
	Amount          int
	Periodicity     string
	TrialPeriodDays int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewPlanFrom(input NewPlanFromInput) *Plan {
	return &Plan{
		id:              input.ID,
		name:            input.Name,
		description:     input.Description,
		amount:          input.Amount,
		periodicity:     input.Periodicity,
		trialPeriodDays: input.TrialPeriodDays,
		createdAt:       input.CreatedAt,
		updatedAt:       input.UpdatedAt,
	}
}

func (p *Plan) ID() string {
	return p.id
}

func (p *Plan) Name() string {
	return p.name
}

func (p *Plan) Description() string {
	return p.description
}

func (p *Plan) Amount() int {
	return p.amount
}

func (p *Plan) Periodicity() string {
	return p.periodicity
}

func (p *Plan) TrialPeriodDays() int {
	return p.trialPeriodDays
}

func (p *Plan) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Plan) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Plan) ValidatePeriodicity() error {
	validIntervals := map[string]string{
		MONTHLY_PLAN_PERIODICITY: MONTHLY_PLAN_PERIODICITY,
		YEARLY_PLAN_PERIODICITY:  YEARLY_PLAN_PERIODICITY,
	}

	if _, ok := validIntervals[p.periodicity]; !ok {
		return core_err.NewModelErr("invalid periodicity")
	}

	return nil
}
