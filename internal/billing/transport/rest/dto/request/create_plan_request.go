package request

type CreatePlanRequest struct {
	Name            string `json:"name" validate:"required"`
	Description     string `json:"description" validate:"required"`
	Amount          int    `json:"amount" validate:"min=0"`
	Periodicity     string `json:"periodicity" validate:"required"`
	TrialPeriodDays int    `json:"trial_period_days" validate:"min=0"`
}
