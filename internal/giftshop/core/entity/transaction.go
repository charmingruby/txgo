package entity

import (
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type NewTransactionInput struct {
	isInPoints    bool
	amountInCents int
	buyerWallet   *Wallet
}

func NewTransaction(in NewTransactionInput) (*Transaction, error) {
	t := Transaction{
		isInPoints:    in.isInPoints,
		buyerWallet:   in.buyerWallet,
		amountInCents: in.amountInCents,
		BaseEntity:    core.NewBaseEntity(),
	}

	if err := t.validate(); err != nil {
		return nil, err
	}

	return &t, nil
}

func NewTransactionFrom(in Transaction) *Transaction {
	return &Transaction{
		isInPoints:    in.isInPoints,
		buyerWallet:   in.buyerWallet,
		amountInCents: in.amountInCents,
		BaseEntity:    in.BaseEntity,
	}
}

func (t *Transaction) validate() error {
	if t.amountInCents <= 0 {
		return core_err.NewEntityErr("amountInCents must be greater than 0")
	}

	if t.buyerWallet == nil {
		return core_err.NewEntityErr("buyerWallet is required")
	}

	return nil
}

type Transaction struct {
	core.BaseEntity

	isInPoints    bool
	amountInCents int
	buyerWallet   *Wallet
}
