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
	Name              string
	Message           string
	BaseValueInPoints int
	SenderEmail       string
	ReceiverEmail     string
}

func NewGift(in NewGiftInput) (*Gift, error) {
	g := Gift{
		name:             in.Name,
		message:          in.Message,
		senderEmail:      in.SenderEmail,
		receiverEmail:    in.ReceiverEmail,
		baseValueInPoins: in.BaseValueInPoints,
		status:           GIFT_STATUS_PENDING,
		payment:          nil,
		BaseEntity:       core.NewBaseEntity(),
	}

	if err := g.validate(); err != nil {
		return nil, err
	}

	return &g, nil
}

func NewGiftFrom(in Gift) *Gift {
	return &Gift{
		name:             in.name,
		message:          in.message,
		receiverEmail:    in.receiverEmail,
		senderEmail:      in.senderEmail,
		baseValueInPoins: in.baseValueInPoins,
		status:           in.status,
		payment:          in.payment,
		BaseEntity:       in.BaseEntity,
	}
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

	if g.baseValueInPoins < 0 {
		return core_err.NewEntityErr("baseValueInPoins should be greater than 0")
	}

	return nil
}

type Gift struct {
	core.BaseEntity

	name             string
	message          string
	receiverEmail    string
	senderEmail      string
	baseValueInPoins int
	status           string
	payment          *Payment
}
