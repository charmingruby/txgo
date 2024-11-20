package service

import (
	"github.com/charmingruby/txgo/internal/giftshop/core/repository"
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/integration"
)

type TransactionalConsistencyParams struct {
	GiftRepository                    repository.GiftRepository
	PaymentRepository                 repository.PaymentRepository
	TransactionRepository             repository.TransactionRepository
	WalletRepository                  repository.WalletRepository
	BillingSubscriptionStatusProvider integration.BillingSubscriptionStatusIntegration
}

type Service struct {
	paymentRepo                       repository.PaymentRepository
	giftRepo                          repository.GiftRepository
	walletRepo                        repository.WalletRepository
	transactionRepo                   repository.TransactionRepository
	transactionalConsistencyProvider  core.TransactionalConsistencyProvider[TransactionalConsistencyParams]
	billingSubscriptionStatusProvider integration.BillingSubscriptionStatusIntegration
}

func New(
	paymentRepo repository.PaymentRepository,
	giftRepo repository.GiftRepository,
	walletRepo repository.WalletRepository,
	transactionRepo repository.TransactionRepository,
	transactionalConsistencyProvider core.TransactionalConsistencyProvider[TransactionalConsistencyParams],
	billingSubscriptionStatusProvider integration.BillingSubscriptionStatusIntegration,
) *Service {
	return &Service{
		paymentRepo:                       paymentRepo,
		giftRepo:                          giftRepo,
		walletRepo:                        walletRepo,
		transactionRepo:                   transactionRepo,
		transactionalConsistencyProvider:  transactionalConsistencyProvider,
		billingSubscriptionStatusProvider: billingSubscriptionStatusProvider,
	}
}
