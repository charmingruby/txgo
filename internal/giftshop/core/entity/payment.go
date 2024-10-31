package entity

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
)

type NewPaymentInput struct {
	method            string
	installments      int
	tax               int
	totalValueInCents int
}

func NewPayment(in NewPaymentInput) (*Payment, error) {
	g := Payment{
		id:                core.NewID(),
		method:            in.method,
		installments:      in.installments,
		tax:               in.tax,
		totalValueInCents: in.totalValueInCents,
		createdAt:         time.Now(),
	}

	return &g, nil
}

type Payment struct {
	id                string
	method            string
	installments      int
	tax               int
	totalValueInCents int
	status            string
	createdAt         time.Time
}
