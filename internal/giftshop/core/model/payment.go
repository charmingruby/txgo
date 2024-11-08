package model

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

const (
	PAYMENT_STATUS_PENDING = "PENDING"
	PAYMENT_STATUS_PAID    = "PAID"
	PAYMENT_STATUS_FAILED  = "FAILED"
)

type NewPaymentInput struct {
	Installments int
	TaxPercent   int
	TotalValue   int
}

func NewPayment(in NewPaymentInput) (*Payment, error) {
	g := Payment{
		id:           core.NewID(),
		installments: in.Installments,
		taxPercent:   in.TaxPercent,
		partialValue: 0,
		totalValue:   in.TotalValue,
		status:       PAYMENT_STATUS_PENDING,
		transaction:  nil,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}

	if err := g.validate(); err != nil {
		return nil, err
	}

	return &g, nil
}

func (p *Payment) validate() error {
	if p.installments < 1 {
		return core_err.NewModelErr("installments must be greater than or equal to 1")
	}

	if p.taxPercent <= 0 {
		return core_err.NewModelErr("taxPercent must be greater than or equal to 0")
	}

	if p.totalValue <= 0 {
		return core_err.NewModelErr("totalValue must be greater than 0")
	}

	return nil
}

type Payment struct {
	id           string
	installments int
	taxPercent   int
	partialValue int
	totalValue   int
	status       string
	transaction  *Transaction
	createdAt    time.Time
	updatedAt    time.Time
}
