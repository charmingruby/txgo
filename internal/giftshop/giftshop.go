package giftshop

import (
	"github.com/charmingruby/txgo/internal/giftshop/transport/rest/endpoint"
	"github.com/go-chi/chi/v5"
)

func NewHTTPHandler(r *chi.Mux) {
	endpoint.New(r).Register()
}
