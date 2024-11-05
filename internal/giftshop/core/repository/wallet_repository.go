package repository

import "github.com/charmingruby/txgo/internal/giftshop/core/entity"

type WalletRepository interface {
	FindByOwnerEmail(ownerEmail string) (*entity.Wallet, error)
	Store(Wallet *entity.Wallet) error
}
