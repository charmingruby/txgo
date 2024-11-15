package repository

import "github.com/charmingruby/txgo/internal/giftshop/core/model"

type PaymentRepository interface {
	FindByID(id string) (*model.Payment, error)
	Store(payment *model.Payment) error
	UpdateTransactionIDAndStatusByID(payment *model.Payment) error
}
