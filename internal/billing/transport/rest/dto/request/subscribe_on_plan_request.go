package request

type SubscribeOnPlanRequest struct {
	Email     string `json:"email" validate:"required,email"`
	AutoRenew bool   `json:"auto_renew"`
}
