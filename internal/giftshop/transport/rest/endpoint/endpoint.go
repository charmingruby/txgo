package endpoint

import (
	"github.com/charmingruby/txgo/internal/giftshop/core/service"
	"github.com/go-chi/chi/v5"
)

type Endpoint struct {
	router  *chi.Mux
	service *service.Service
}

func New(r *chi.Mux) *Endpoint {
	return &Endpoint{router: r}
}

func (e *Endpoint) Register() {
	e.router.Post("/wallets", e.createWalletHandler())

	e.router.Route("/gifts", func(r chi.Router) {
		r.Post("/customize", e.giftCustomizationHandler())
		r.Post("/checkout", e.giftCheckoutHandler())
	})
}
