package entity

import (
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
		installments: in.Installments,
		taxPercent:   in.TaxPercent,
		partialValue: 0,
		totalValue:   in.TotalValue,
		status:       PAYMENT_STATUS_PENDING,
		transaction:  nil,
		BaseEntity:   core.NewBaseEntity(),
	}

	if err := g.validate(); err != nil {
		return nil, err
	}

	return &g, nil
}

func (p *Payment) validate() error {
	if p.installments < 1 {
		return core_err.NewEntityErr("installments must be greater than or equal to 1")
	}

	if p.taxPercent <= 0 {
		return core_err.NewEntityErr("taxPercent must be greater than or equal to 0")
	}

	if p.totalValue <= 0 {
		return core_err.NewEntityErr("totalValue must be greater than 0")
	}

	return nil
}

func (p *Payment) CalculatePartialValue() {
	totalValueWithTax := p.totalValue + (p.totalValue * p.taxPercent / 100)
	p.partialValue = totalValueWithTax / p.installments
}

type Payment struct {
	core.BaseEntity

	installments int
	taxPercent   int
	partialValue int
	totalValue   int
	status       string
	transaction  *Transaction
}
