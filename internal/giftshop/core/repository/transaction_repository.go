package repository

import "github.com/charmingruby/txgo/internal/giftshop/core/model"

type TransactionRepository interface {
	FindByID(id string) (*model.Transaction, error)
	Store(transaction *model.Transaction) error
}
