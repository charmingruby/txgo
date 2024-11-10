package model

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type Transaction struct {
	id               string
	points           int
	receiverWalletID string
	senderWalletID   string
	createdAt        time.Time
	updatedAt        time.Time
}

type NewTransactionInput struct {
	Points           int
	ReceiverWalletID string
	SenderWalletID   string
}

func NewTransaction(in NewTransactionInput) (*Transaction, error) {
	t := Transaction{
		id:               core.NewID(),
		points:           in.Points,
		receiverWalletID: in.ReceiverWalletID,
		senderWalletID:   in.SenderWalletID,
		createdAt:        time.Now(),
		updatedAt:        time.Now(),
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

func (t *Transaction) ID() string {
	return t.id
}

func (t *Transaction) Points() int {
	return t.points
}

func (t *Transaction) ReceiverWalletID() string {
	return t.receiverWalletID
}

func (t *Transaction) SenderWalletID() string {
	return t.senderWalletID
}

func (t *Transaction) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Transaction) UpdatedAt() time.Time {
	return t.updatedAt
}
