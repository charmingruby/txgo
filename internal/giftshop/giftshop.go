package giftshop

import (
	"database/sql"

	"github.com/charmingruby/txgo/internal/giftshop/core/repository"
	"github.com/charmingruby/txgo/internal/giftshop/core/service"
	"github.com/charmingruby/txgo/internal/giftshop/database/mysql"
	"github.com/charmingruby/txgo/internal/giftshop/transport/rest/endpoint"
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/integration"
	"github.com/go-chi/chi/v5"
)

func NewService(
	walletRepository repository.WalletRepository,
	giftRepository repository.GiftRepository,
	paymentRepository repository.PaymentRepository,
	transactionRepository repository.TransactionRepository,
	transactionalConsistencyProvider core.TransactionalConsistencyProvider[service.TransactionalConsistencyParams],
	billingSubscriptionStatusProvider integration.BillingSubscriptionStatusIntegration,

) *service.Service {
	return service.New(
		paymentRepository,
		giftRepository,
		walletRepository,
		transactionRepository,
		transactionalConsistencyProvider,
		billingSubscriptionStatusProvider,
	)
}

func NewWalletRepository(db *sql.DB) repository.WalletRepository {
	return mysql.NewWalletRepository(db)
}

func NewGiftRepository(db *sql.DB) repository.GiftRepository {
	return mysql.NewGiftRepository(db)
}

func NewPaymentRepository(db *sql.DB) repository.PaymentRepository {
	return mysql.NewPaymentRepository(db)
}

func NewTransactionRepository(db *sql.DB) repository.TransactionRepository {
	return mysql.NewTransactionRepository(db)
}

func NewTransactionConsistencyProvider(db *sql.DB) *mysql.TransactionConsistencyProvider {
	return mysql.NewTransactionConsistencyProvider(db)
}

func NewHTTPHandler(r *chi.Mux, service *service.Service) {
	endpoint.New(r, service).Register()
}
