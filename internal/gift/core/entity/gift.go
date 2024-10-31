package entity

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type NewGiftInput struct {
	name          string
	message       string
	senderEmail   string
	receiverEmail string
	amountInCents int
}

func NewGift(in NewGiftInput) (*Gift, error) {
	g := Gift{
		id:            core.NewID(),
		name:          in.name,
		message:       in.message,
		senderEmail:   in.senderEmail,
		receiverEmail: in.receiverEmail,
		amountInCents: in.amountInCents,
		createdAt:     time.Now(),
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

	if g.amountInCents <= 0 {
		return core_err.NewEntityErr("amount must be greater than 0")
	}

	return nil
}

type Gift struct {
	id            string
	name          string
	message       string
	receiverEmail string
	senderEmail   string
	amountInCents int
	createdAt     time.Time
}
