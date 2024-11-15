package model

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type Transaction struct {
	id            string
	points        int
	payerWalletID string
	createdAt     time.Time
	updatedAt     time.Time
}

type NewTransactionInput struct {
	Points        int
	PayerWalletID string
}

func NewTransaction(in NewTransactionInput) (*Transaction, error) {
	t := Transaction{
		id:            core.NewID(),
		points:        in.Points,
		payerWalletID: in.PayerWalletID,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}

	if err := t.validate(); err != nil {
		return nil, err
	}

	return &t, nil
}

type NewTransactionFromInput struct {
	ID            string
	Points        int
	PayerWalletID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewTransactionFrom(in NewTransactionFromInput) *Transaction {
	return &Transaction{
		id:            in.ID,
		points:        in.Points,
		payerWalletID: in.PayerWalletID,
		createdAt:     in.CreatedAt,
		updatedAt:     in.UpdatedAt,
	}
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

func (t *Transaction) PayerWalletID() string {
	return t.payerWalletID
}

func (t *Transaction) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Transaction) UpdatedAt() time.Time {
	return t.updatedAt
}
