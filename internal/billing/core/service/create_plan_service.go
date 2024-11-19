package service

import (
	"github.com/charmingruby/txgo/internal/billing/core/model"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type CreatePlanInput struct {
	Name            string
	Description     string
	Amount          int
	Periodicity     string
	TrialPeriodDays int
}

func (s *Service) CreatePlanService(params CreatePlanInput) error {
	planFound, err := s.planRepo.FindByName(params.Name)
	if err != nil {
		return err
	}

	if planFound != nil {
		return core_err.NewResourceAlreadyExistsErr("plan")
	}

	plan := model.NewPlan(model.NewPlanInput{
		Name:            params.Name,
		Description:     params.Description,
		Amount:          params.Amount,
		Periodicity:     params.Periodicity,
		TrialPeriodDays: params.TrialPeriodDays,
	})

	if err := plan.ValidatePeriodicity(); err != nil {
		return err
	}

	if err := s.planRepo.Store(plan); err != nil {
		return err
	}

	return nil
}
