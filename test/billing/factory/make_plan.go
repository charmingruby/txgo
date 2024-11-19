package factory

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/charmingruby/txgo/internal/billing/core/model"
	"github.com/charmingruby/txgo/internal/billing/core/repository"
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/helper"
)

func MakePlan(planRepo repository.PlanRepository, params model.NewPlanFromInput) (*model.Plan, error) {
	plan := createPlan(params)

	if err := planRepo.Store(plan); err != nil {
		return nil, err
	}

	return plan, nil
}

func createPlan(params model.NewPlanFromInput) *model.Plan {
	input := model.NewPlanFromInput{
		ID:              helper.If[string](params.ID != "", params.ID, core.NewID()),
		Name:            helper.If[string](params.Name != "", params.Name, gofakeit.Name()),
		Description:     helper.If[string](params.Description != "", params.Description, gofakeit.Sentence(10)),
		Amount:          params.Amount,
		Periodicity:     helper.If[string](params.Periodicity != "", params.Periodicity, model.MONTHLY_PLAN_PERIODICITY),
		TrialPeriodDays: params.TrialPeriodDays,
		CreatedAt:       helper.If[time.Time](params.CreatedAt != time.Time{}, params.CreatedAt, time.Now()),
		UpdatedAt:       helper.If[time.Time](params.UpdatedAt != time.Time{}, params.UpdatedAt, time.Now()),
	}

	return model.NewPlanFrom(input)
}
