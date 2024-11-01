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
	installments     int
	taxPercent       int
	totalValuePoints int
}

func NewPayment(in NewPaymentInput) (*Payment, error) {
	g := Payment{
		installments:       in.installments,
		taxPercent:         in.taxPercent,
		partialValuePoints: 0,
		totalValuePoints:   in.totalValuePoints,
		status:             PAYMENT_STATUS_PENDING,
		transaction:        nil,
		BaseEntity:         core.NewBaseEntity(),
	}

	if err := g.validate(); err != nil {
		return nil, err
	}

	return &g, nil
}

func NewPaymentFrom(in Payment) *Payment {
	return &Payment{
		installments:       in.installments,
		taxPercent:         in.taxPercent,
		partialValuePoints: in.partialValuePoints,
		totalValuePoints:   in.totalValuePoints,
		status:             PAYMENT_STATUS_PENDING,
		transaction:        in.transaction,
		BaseEntity:         in.BaseEntity,
	}
}

func (p *Payment) validate() error {
	if p.installments < 1 {
		return core_err.NewEntityErr("installments must be greater than or equal to 1")
	}

	if p.taxPercent <= 0 {
		return core_err.NewEntityErr("taxPercent must be greater than or equal to 0")
	}

	if p.totalValuePoints <= 0 {
		return core_err.NewEntityErr("totalValuePoints must be greater than 0")
	}

	return nil
}

func (p *Payment) CalculatePartialValue() {
	totalValueWithTax := p.totalValuePoints + (p.totalValuePoints * p.taxPercent / 100)
	p.partialValuePoints = totalValueWithTax / p.installments
}

type Payment struct {
	core.BaseEntity

	installments       int
	taxPercent         int
	partialValuePoints int
	totalValuePoints   int
	status             string
	transaction        *Transaction
}
