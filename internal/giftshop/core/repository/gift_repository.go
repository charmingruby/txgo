package repository

import "github.com/charmingruby/txgo/internal/giftshop/core/entity"

type GiftRepository interface {
	FindByID(id string) (*entity.Gift, error)
	Store(gift *entity.Gift) error
}
