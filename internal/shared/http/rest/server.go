package rest

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type RestServer struct {
	httpServer *http.Server
	router     *chi.Mux
}

func NewServer(port string, router *chi.Mux) *RestServer {
	httpServer := http.Server{
		Addr: ":" + port,
	}

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	return &RestServer{
		httpServer: &httpServer,
		router:     router,
	}
}

func (s *RestServer) Run() error {
	if err := http.ListenAndServe(":3000", s.router); err != nil {
		return err
	}

	return nil
}
