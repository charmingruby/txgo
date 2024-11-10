package repository

import "github.com/charmingruby/txgo/internal/giftshop/core/model"

type PaymentRepository interface {
	Store(payment *model.Payment) error
	UpdateTransactionIDAndStatusByID(payment *model.Payment) error
}
