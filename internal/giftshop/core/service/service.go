package service

import (
	"github.com/charmingruby/txgo/internal/giftshop/core/repository"
	"github.com/charmingruby/txgo/internal/shared/core"
)

type TransactionalConsistencyParams struct {
	GiftRepository        repository.GiftRepository
	PaymentRepository     repository.PaymentRepository
	TransactionRepository repository.TransactionRepository
	WalletRepository      repository.WalletRepository
}

type Service struct {
	paymentRepo                      repository.PaymentRepository
	giftRepo                         repository.GiftRepository
	walletRepo                       repository.WalletRepository
	transactionRepo                  repository.TransactionRepository
	transactionalConsistencyProvider core.TransactionalConsistencyProvider[TransactionalConsistencyParams]
}

func New(
	paymentRepo repository.PaymentRepository,
	giftRepo repository.GiftRepository,
	walletRepo repository.WalletRepository,
	transactionRepo repository.TransactionRepository,
	transactionalConsistencyProvider core.TransactionalConsistencyProvider[TransactionalConsistencyParams],
) *Service {
	return &Service{
		paymentRepo:                      paymentRepo,
		giftRepo:                         giftRepo,
		walletRepo:                       walletRepo,
		transactionRepo:                  transactionRepo,
		transactionalConsistencyProvider: transactionalConsistencyProvider,
	}
}
