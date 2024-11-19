package model

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
)

const (
	MONTHLY_PLAN_INTERVAL = "month"
	YEARLY_PLAN_INTERVAL  = "year"
)

type Plan struct {
	id              string
	name            string
	description     string
	amount          int
	interval        string
	trialPeriodDays int
	createdAt       time.Time
	updateAt        time.Time
}

type NewPlanInput struct {
	Name            string
	Description     string
	Amount          int
	Interval        string
	TrialPeriodDays int
}

func NewPlan(input NewPlanInput) *Plan {
	return &Plan{
		id:              core.NewID(),
		name:            input.Name,
		description:     input.Description,
		amount:          input.Amount,
		interval:        input.Interval,
		trialPeriodDays: input.TrialPeriodDays,
		createdAt:       time.Now(),
		updateAt:        time.Now(),
	}
}
