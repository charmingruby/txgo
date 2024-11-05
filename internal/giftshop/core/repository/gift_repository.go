package repository

import "github.com/charmingruby/txgo/internal/giftshop/core/entity"

type GiftRepository interface {
	Store(gift *entity.Gift) error
}
