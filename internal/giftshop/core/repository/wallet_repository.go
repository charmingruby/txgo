package repository

import "github.com/charmingruby/txgo/internal/giftshop/core/model"

type WalletRepository interface {
	FindByOwnerEmail(ownerEmail string) (*model.Wallet, error)
	Store(Wallet *model.Wallet) error
}
