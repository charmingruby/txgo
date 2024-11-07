package entity

import (
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

const (
	GIFT_STATUS_PENDING  = "PENDING"
	GIFT_STATUS_SENT     = "SENT"
	GIFT_STATUS_ACCEPTED = "ACCEPTED"
)

type NewGiftInput struct {
	Name          string
	Message       string
	BaseValue     int
	SenderEmail   string
	ReceiverEmail string
}

func NewGift(in NewGiftInput) (*Gift, error) {
	g := Gift{
		name:          in.Name,
		message:       in.Message,
		senderEmail:   in.SenderEmail,
		receiverEmail: in.ReceiverEmail,
		baseValue:     in.BaseValue,
		status:        GIFT_STATUS_PENDING,
		payment:       nil,
		BaseEntity:    core.NewBaseEntity(),
	}

	if err := g.validate(); err != nil {
		return nil, err
	}

	return &g, nil
}

func (g *Gift) validate() error {
	if g.name == "" {
		return core_err.NewEntityErr("name is required")
	}

	if g.senderEmail == "" {
		return core_err.NewEntityErr("senderEmail is required")
	}

	if g.receiverEmail == "" {
		return core_err.NewEntityErr("receiverEmail is required")
	}

	if g.baseValue < 0 {
		return core_err.NewEntityErr("baseValue should be greater than 0")
	}

	return nil
}

func (g *Gift) SenderEmail() string {
	return g.senderEmail
}

func (g *Gift) ReceiverEmail() string {
	return g.receiverEmail
}

func (g *Gift) BaseValue() int {
	return g.baseValue
}

type Gift struct {
	core.BaseEntity

	name          string
	message       string
	receiverEmail string
	senderEmail   string
	baseValue     int
	status        string
	payment       *Payment
}
