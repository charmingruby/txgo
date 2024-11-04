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

func NewServer(port string) *RestServer {
	router := chi.NewRouter()

	httpServer := http.Server{
		Addr: ":" + port,
	}

	return &RestServer{
		httpServer: &httpServer,
		router:     router,
	}
}

func (s *RestServer) Run() error {
	s.router.Use(middleware.Logger)

	if err := http.ListenAndServe(":3000", s.router); err != nil {
		return err
	}

	return nil
}
