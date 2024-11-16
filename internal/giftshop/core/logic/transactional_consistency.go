package logic

import "github.com/charmingruby/txgo/internal/giftshop/core/repository"

type TransactionalConsistencyLogicParams struct {
	giftRepository        repository.GiftRepository
	paymentRepository     repository.PaymentRepository
	transactionRepository repository.TransactionRepository
	walletRepository      repository.WalletRepository
}

type TransactionalConsistencyLogicProvider interface {
	Transact(func(context TransactionalConsistencyLogicParams) error) error
}
