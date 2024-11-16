package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type RestServer struct {
	HttpServer *http.Server
	Router     *chi.Mux
}

func NewServer(port string, router *chi.Mux) *RestServer {
	httpServer := http.Server{
		Addr: ":" + port,
	}

	attachBaseMiddlewares(router)

	return &RestServer{
		HttpServer: &httpServer,
		Router:     router,
	}
}

func (s *RestServer) Run() error {
	if err := http.ListenAndServe(s.HttpServer.Addr, s.Router); err != nil {
		return err
	}

	return nil
}

func (s *RestServer) Shutdown(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}

func attachBaseMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	router.Use(cors.Handler)
}
