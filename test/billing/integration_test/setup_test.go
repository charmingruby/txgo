package integration_test

import (
	"net/http/httptest"
	"testing"

	"github.com/charmingruby/txgo/internal/billing/core/repository"
	"github.com/charmingruby/txgo/internal/billing/core/service"
	"github.com/charmingruby/txgo/internal/billing/database/mysql"
	"github.com/charmingruby/txgo/internal/billing/transport/rest/endpoint"
	"github.com/charmingruby/txgo/internal/shared/transport/rest"
	"github.com/charmingruby/txgo/test/shared/container"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite

	dbContainer      *container.MySQL
	server           *httptest.Server
	handler          *endpoint.Endpoint
	planRepo         repository.PlanRepository
	subscriptionRepo repository.SubscriptionRepository
}

func (s *Suite) SetupSuite() {
	s.dbContainer = container.NewMySQL()
}

func (s *Suite) TearDownSuite() {
	err := s.dbContainer.Teardown()
	s.NoError(err)
}

func (s *Suite) SetupTest() {
	router := chi.NewRouter()

	s.planRepo = mysql.NewPlanRepository(s.dbContainer.DB)
	s.subscriptionRepo = mysql.NewSubscriptionRepository(s.dbContainer.DB)

	service := service.New(
		s.subscriptionRepo,
		s.planRepo,
	)

	server := rest.NewServer("3000", router)

	s.handler = endpoint.New(router, service)
	s.handler.Register()

	s.server = httptest.NewServer(server.Router)
}

func (s *Suite) SetupSubTest() {
	err := s.dbContainer.RunMigrations()
	s.NoError(err)
}

func (s *Suite) TearDownSubTest() {
	err := s.dbContainer.RollbackMigrations()
	s.NoError(err)
}

func (s *Suite) TearDownTest() {
	s.server.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
