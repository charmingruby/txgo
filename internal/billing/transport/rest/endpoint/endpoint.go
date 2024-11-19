package endpoint

import (
	"github.com/charmingruby/txgo/internal/billing/core/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Endpoint struct {
	router    *chi.Mux
	service   *service.Service
	validator *validator.Validate
}

func New(r *chi.Mux, service *service.Service) *Endpoint {
	return &Endpoint{
		router:    r,
		service:   service,
		validator: validator.New(),
	}
}

func (e *Endpoint) Register() {
	e.router.Post("/plans", e.createPlanHandler())
	e.router.Post("/plans/{planID}/subscribe", e.subscribeOnPlanHandler())
}
