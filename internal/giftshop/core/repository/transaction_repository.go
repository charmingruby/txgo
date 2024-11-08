package repository

import "github.com/charmingruby/txgo/internal/giftshop/core/model"

type TransactionRepository interface {
	Store(transaction *model.Transaction) error
}
