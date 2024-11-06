package entity

import (
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type NewTransactionInput struct {
	Points         int
	ReceiverWallet *Wallet
	SenderWallet   *Wallet
}

func NewTransaction(in NewTransactionInput) (*Transaction, error) {
	t := Transaction{
		points:         in.Points,
		receiverWallet: in.ReceiverWallet,
		senderWallet:   in.SenderWallet,
		BaseEntity:     core.NewBaseEntity(),
	}

	if err := t.validate(); err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Transaction) validate() error {
	if t.points <= 0 {
		return core_err.NewEntityErr("points must be greater than 0")
	}

	return nil
}

type Transaction struct {
	core.BaseEntity

	points         int
	receiverWallet *Wallet
	senderWallet   *Wallet
}
