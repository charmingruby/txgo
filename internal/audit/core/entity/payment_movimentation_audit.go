package entity

import (
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type NewPaymentMovimentationAuditInput struct {
	actorEmail    string
	context       string
	amountInCents int
}

func NewPaymentMovimentationAudit(in NewPaymentMovimentationAuditInput) (*PaymentMovimentationAudit, error) {
	p := PaymentMovimentationAudit{
		actorEmail:    in.actorEmail,
		context:       in.context,
		amountInCents: in.amountInCents,
		BaseEntity:    core.NewBaseEntity(),
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	return &p, nil
}

func NewPaymentMovimentationAuditFrom(in PaymentMovimentationAudit) *PaymentMovimentationAudit {
	return &PaymentMovimentationAudit{
		actorEmail:    in.actorEmail,
		context:       in.context,
		amountInCents: in.amountInCents,
		BaseEntity:    in.BaseEntity,
	}
}

func (p *PaymentMovimentationAudit) validate() error {
	if p.actorEmail == "" {
		return core_err.NewEntityErr("actorEmail is required")
	}

	if p.context == "" {
		return core_err.NewEntityErr("context is required")
	}

	if p.amountInCents <= 0 {
		return core_err.NewEntityErr("amountInCents must be greater than 0")
	}

	return nil
}

type PaymentMovimentationAudit struct {
	core.BaseEntity

	actorEmail    string
	context       string
	amountInCents int
}
