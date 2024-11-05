package repository

import "github.com/charmingruby/txgo/internal/giftshop/core/entity"

type PaymentRepository interface {
	Store(payment *entity.Payment) error
}
