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

type Payment struct {
	id            string
	installments  int
	taxPercent    int
	partialValue  int
	totalValue    int
	status        string
	transactionID string
	createdAt     time.Time
	updatedAt     time.Time
}

type NewPaymentInput struct {
	Installments int
	TaxPercent   int
	TotalValue   int
}

func NewPayment(in NewPaymentInput) (*Payment, error) {
	g := Payment{
		id:            core.NewID(),
		installments:  in.Installments,
		taxPercent:    in.TaxPercent,
		partialValue:  0,
		totalValue:    in.TotalValue,
		status:        PAYMENT_STATUS_PENDING,
		transactionID: "",
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}

	if err := g.validate(); err != nil {
		return nil, err
	}

	return &g, nil
}

type NewPaymentFromInput struct {
	ID            string
	Installments  int
	TaxPercent    int
	PartialValue  int
	TotalValue    int
	Status        string
	TransactionID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewPaymentFrom(in NewPaymentFromInput) *Payment {
	return &Payment{
		id:            in.ID,
		installments:  in.Installments,
		taxPercent:    in.TaxPercent,
		partialValue:  in.PartialValue,
		totalValue:    in.TotalValue,
		status:        in.Status,
		transactionID: in.TransactionID,
		createdAt:     in.CreatedAt,
		updatedAt:     in.UpdatedAt,
	}
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

func (p *Payment) ID() string {
	return p.id
}

func (p *Payment) Installments() int {
	return p.installments
}

func (p *Payment) TaxPercent() int {
	return p.taxPercent
}

func (p *Payment) PartialValue() int {
	return p.partialValue
}

func (p *Payment) SetPartialValue(value int) {
	p.touch()
	p.partialValue = value
}

func (p *Payment) TotalValue() int {
	return p.totalValue
}

func (p *Payment) Status() string {
	return p.status
}

func (p *Payment) Paid() {
	p.touch()
	p.status = PAYMENT_STATUS_PAID
}

func (p *Payment) TransactionID() string {
	return p.transactionID
}

func (p *Payment) SetTransactionID(transactionID string) {
	p.touch()
	p.transactionID = transactionID
}

func (p *Payment) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Payment) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Payment) touch() {
	p.updatedAt = time.Now()
}
