package repository

import "github.com/charmingruby/txgo/internal/giftshop/core/model"

type WalletRepository interface {
	FindByID(id string) (*model.Wallet, error)
	FindByOwnerEmail(ownerEmail string) (*model.Wallet, error)
	Store(wallet *model.Wallet) error
	UpdatePointsByID(wallet *model.Wallet) error
}
