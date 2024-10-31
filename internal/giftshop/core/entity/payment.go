package entity

import (
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

const (
	PAYMENT_METHOD_CREDIT_CARD = "CREDIT_CARD"
	PAYMENT_METHOD_CASH        = "CASH"
	PAYMENT_METHOD_PIX         = "PIX"
	PAYMENT_METHOD_POINTS      = "POINTS"

	PAYMENT_STATUS_PENDING = "PENDING"
	PAYMENT_STATUS_PAID    = "PAID"
	PAYMENT_STATUS_FAILED  = "FAILED"

	PAYMENT_POINTS_TO_CURRENCY_CENTS = 100
)

type NewPaymentInput struct {
	method            string
	installments      int
	taxPercent        int
	totalValueInCents int
	appliedPoints     int
}

func NewPayment(in NewPaymentInput) (*Payment, error) {
	g := Payment{
		method:              in.method,
		installments:        in.installments,
		taxPercent:          in.taxPercent,
		partialValueInCents: 0,
		totalValueInCents:   in.totalValueInCents,
		appliedPoints:       in.appliedPoints,
		status:              PAYMENT_STATUS_PENDING,
		BaseEntity:          core.NewBaseEntity(),
	}

	if err := g.validate(); err != nil {
		return nil, err
	}

	return &g, nil
}

func NewPaymentFrom(in Payment) *Payment {
	return &Payment{
		method:              in.method,
		installments:        in.installments,
		taxPercent:          in.taxPercent,
		partialValueInCents: in.partialValueInCents,
		totalValueInCents:   in.totalValueInCents,
		appliedPoints:       in.appliedPoints,
		status:              PAYMENT_STATUS_PENDING,
		BaseEntity:          in.BaseEntity,
	}
}

func (p *Payment) validate() error {
	if !p.isPaymentMethodValid() {
		return core_err.NewEntityErr("invalid payment method")
	}

	if p.installments < 1 {
		return core_err.NewEntityErr("installments must be greater than or equal to 1")
	}

	if p.taxPercent < 0 {
		return core_err.NewEntityErr("taxPercent must be greater than or equal to 0")
	}

	if p.totalValueInCents < 0 {
		return core_err.NewEntityErr("totalValueInCents must be greater than or equal to 0")
	}

	if p.appliedPoints < 0 {
		return core_err.NewEntityErr("appliedPoints must be greater than or equal to 0")
	}

	return nil
}

func (p *Payment) CalculatePartialValue() error {
	totalValueWithTax := p.totalValueInCents + (p.totalValueInCents * p.taxPercent / 100)

	var appliedPointsInCents int
	hasAppliedPoints := p.appliedPoints > 0
	if hasAppliedPoints {
		appliedPointsInCents = p.appliedPoints * PAYMENT_POINTS_TO_CURRENCY_CENTS

		if appliedPointsInCents > totalValueWithTax {
			return core_err.NewEntityErr("appliedPoints cannot be greater than totalValueInCents")
		}

		if appliedPointsInCents == totalValueWithTax {
			p.partialValueInCents = 0
			p.method = PAYMENT_METHOD_POINTS
			return nil
		}
	}

	totalValueWithTax -= appliedPointsInCents

	p.partialValueInCents = totalValueWithTax / p.installments

	return nil
}

func (p *Payment) isPaymentMethodValid() bool {
	methods := map[string]bool{
		PAYMENT_METHOD_CREDIT_CARD: true,
		PAYMENT_METHOD_CASH:        true,
		PAYMENT_METHOD_PIX:         true,
		PAYMENT_METHOD_POINTS:      true,
	}

	return methods[p.method]
}

type Payment struct {
	core.BaseEntity

	method              string
	installments        int
	taxPercent          int
	partialValueInCents int
	totalValueInCents   int
	appliedPoints       int
	status              string
}
