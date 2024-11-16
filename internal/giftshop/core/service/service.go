package service

import (
	"github.com/charmingruby/txgo/internal/giftshop/core/logic"
	"github.com/charmingruby/txgo/internal/giftshop/core/repository"
)

type Service struct {
	paymentRepo                      repository.PaymentRepository
	giftRepo                         repository.GiftRepository
	walletRepo                       repository.WalletRepository
	transactionRepo                  repository.TransactionRepository
	transactionalConsistencyProvider logic.TransactionalConsistencyProvider
}

func New(
	paymentRepo repository.PaymentRepository,
	giftRepo repository.GiftRepository,
	walletRepo repository.WalletRepository,
	transactionRepo repository.TransactionRepository,
	transactionalConsistencyProvider logic.TransactionalConsistencyProvider,
) *Service {
	return &Service{
		paymentRepo:                      paymentRepo,
		giftRepo:                         giftRepo,
		walletRepo:                       walletRepo,
		transactionRepo:                  transactionRepo,
		transactionalConsistencyProvider: transactionalConsistencyProvider,
	}
}
