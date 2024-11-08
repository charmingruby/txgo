package model

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type Transaction struct {
	id             string
	points         int
	receiverWallet *Wallet
	senderWallet   *Wallet
	createdAt      time.Time
	updatedAt      time.Time
}

type NewTransactionInput struct {
	Points         int
	ReceiverWallet *Wallet
	SenderWallet   *Wallet
}

func NewTransaction(in NewTransactionInput) (*Transaction, error) {
	t := Transaction{
		id:             core.NewID(),
		points:         in.Points,
		receiverWallet: in.ReceiverWallet,
		senderWallet:   in.SenderWallet,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
	}

	if err := t.validate(); err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Transaction) validate() error {
	if t.points <= 0 {
		return core_err.NewModelErr("points must be greater than 0")
	}

	return nil
}
