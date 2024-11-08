package repository

import "github.com/charmingruby/txgo/internal/giftshop/core/model"

type GiftRepository interface {
	FindByID(id string) (*model.Gift, error)
	Store(gift *model.Gift) error
}
