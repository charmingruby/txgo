package entity

import (
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type NewTransactionInput struct {
	AmountInPoints int
	ReceiverWallet *Wallet
	BuyerWallet    *Wallet
}

func NewTransaction(in NewTransactionInput) (*Transaction, error) {
	t := Transaction{
		amountInPoints: in.AmountInPoints,
		receiverWallet: in.ReceiverWallet,
		buyerWallet:    in.BuyerWallet,
		BaseEntity:     core.NewBaseEntity(),
	}

	if err := t.validate(); err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Transaction) validate() error {
	if t.amountInPoints <= 0 {
		return core_err.NewEntityErr("amountInPoints must be greater than 0")
	}

	return nil
}

type Transaction struct {
	core.BaseEntity

	amountInPoints int
	receiverWallet *Wallet
	buyerWallet    *Wallet
}
