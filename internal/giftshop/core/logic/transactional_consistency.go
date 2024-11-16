package logic

import "github.com/charmingruby/txgo/internal/giftshop/core/repository"

type TransactionalConsistencyParams struct {
	GiftRepository        repository.GiftRepository
	PaymentRepository     repository.PaymentRepository
	TransactionRepository repository.TransactionRepository
	WalletRepository      repository.WalletRepository
}

type TransactionalConsistencyProvider interface {
	Transact(func(context TransactionalConsistencyParams) error) error
}
