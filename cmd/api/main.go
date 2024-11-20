package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmingruby/txgo/config"
	"github.com/charmingruby/txgo/internal/billing"
	"github.com/charmingruby/txgo/internal/giftshop"
	"github.com/charmingruby/txgo/internal/shared/transport/rest"
	"github.com/charmingruby/txgo/pkg/mysql"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
	}

	config, err := config.New()
	if err != nil {
		slog.Error(fmt.Sprintf("CONFIGURATION: %v", err))
		os.Exit(1)
	}

	db, err := mysql.New(mysql.MySQLConnectionInput{
		Username:     config.MySQLConfig.User,
		Password:     config.MySQLConfig.Password,
		Host:         config.MySQLConfig.Host,
		Port:         config.MySQLConfig.Port,
		DatabaseName: config.MySQLConfig.DatabaseName,
	})
	if err != nil {
		slog.Error(fmt.Sprintf("MYSQL: %v", err))
		os.Exit(1)
	}
	defer db.Close()

	router := chi.NewRouter()

	restServer := rest.NewServer(config.ServerConfig.Port, router)

	initDependencies(router, db)

	shutdown := make(chan error)
	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		slog.Info(fmt.Sprintf("SHUTDOWN: signal caught %s", s))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		slog.Info("SHUTDOWN: Initiating graceful shutdown")
		shutdown <- restServer.Shutdown(ctx)
	}()

	slog.Info(fmt.Sprintf("REST SERVER: Running on port %s", config.ServerConfig.Port))
	if err := restServer.Run(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error(fmt.Sprintf("REST SERVER: %v", err))
			os.Exit(1)
		}
	}

	err = <-shutdown
	if err != nil {
		slog.Error(fmt.Sprintf("REST SERVER: %v", err))
		os.Exit(1)
	}

	slog.Info("REST SERVER: has gracefully shutdown")
}

func initDependencies(r *chi.Mux, db *sql.DB) {
	planRepository := billing.NewPlanRepository(db)
	subscriptionRepository := billing.NewSubscriptionRepository(db)

	billingPublicProvider := billing.NewPublicProvider(subscriptionRepository)
	billingSvc := billing.NewService(planRepository, subscriptionRepository)
	billing.NewHTTPHandler(r, billingSvc)

	walletRepository := giftshop.NewWalletRepository(db)
	giftRepository := giftshop.NewGiftRepository(db)
	transactionRepository := giftshop.NewTransactionRepository(db)
	paymentRepository := giftshop.NewPaymentRepository(db)
	transactionalConsistencyProvider := giftshop.NewTransactionConsistencyProvider(db)

	giftshopSvc := giftshop.NewService(walletRepository, giftRepository, paymentRepository, transactionRepository, transactionalConsistencyProvider, &billingPublicProvider)
	giftshop.NewHTTPHandler(r, giftshopSvc)
}
