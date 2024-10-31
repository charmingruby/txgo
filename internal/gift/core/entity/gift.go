package entity

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

const (
	GIFT_STATUS_PENDING  = "PENDING"
	GIFT_STATUS_SENT     = "SENT"
	GIFT_STATUS_ACCEPTED = "ACCEPTED"
)

type NewGiftInput struct {
	name          string
	message       string
	senderEmail   string
	receiverEmail string
}

func NewGift(in NewGiftInput) (*Gift, error) {
	g := Gift{
		id:            core.NewID(),
		name:          in.name,
		message:       in.message,
		senderEmail:   in.senderEmail,
		receiverEmail: in.receiverEmail,
		status:        GIFT_STATUS_PENDING,
		payment:       nil,
		createdAt:     time.Now(),
	}

	if err := g.validate(); err != nil {
		return nil, err
	}

	return &g, nil
}

func NewGiftFrom(in Gift) *Gift {
	return &Gift{
		id:            in.id,
		name:          in.name,
		message:       in.message,
		receiverEmail: in.receiverEmail,
		senderEmail:   in.senderEmail,
		status:        in.status,
		payment:       in.payment,
		createdAt:     in.createdAt,
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

	return nil
}

type Gift struct {
	id            string
	name          string
	message       string
	receiverEmail string
	senderEmail   string
	status        string
	payment       *Payment
	createdAt     time.Time
}
