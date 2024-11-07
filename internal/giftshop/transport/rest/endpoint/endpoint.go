package endpoint

import (
	"github.com/charmingruby/txgo/internal/giftshop/core/service"
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
	e.router.Post("/wallets", e.createWalletHandler())

	e.router.Route("/gifts", func(r chi.Router) {
		r.Post("/customize", e.giftCustomizationHandler())
		r.Post("/checkout", e.giftCheckoutHandler())
	})
}
