package repository

import "github.com/charmingruby/txgo/internal/giftshop/core/entity"

type TransactionRepository interface {
	Store(transaction *entity.Transaction) error
}
