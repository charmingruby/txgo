package billing

import (
	"database/sql"

	"github.com/charmingruby/txgo/internal/billing/core/repository"
	"github.com/charmingruby/txgo/internal/billing/core/service"
	"github.com/charmingruby/txgo/internal/billing/database/mysql"
	"github.com/charmingruby/txgo/internal/billing/integration/provider"
	"github.com/charmingruby/txgo/internal/billing/transport/rest/endpoint"
	"github.com/go-chi/chi/v5"
)

func NewService(planRepo repository.PlanRepository, subscriptionRepo repository.SubscriptionRepository) *service.Service {
	return service.New(subscriptionRepo, planRepo)
}

func NewPlanRepository(db *sql.DB) repository.PlanRepository {
	return mysql.NewPlanRepository(db)
}

func NewSubscriptionRepository(db *sql.DB) repository.SubscriptionRepository {
	return mysql.NewSubscriptionRepository(db)
}

func NewPublicProvider(subscriptionRepo repository.SubscriptionRepository) provider.Public {
	return *provider.NewPublic(subscriptionRepo)
}

func NewHTTPHandler(r *chi.Mux, service *service.Service) {
	endpoint.New(r, service).Register()
}
