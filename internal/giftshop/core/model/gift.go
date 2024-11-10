package model

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

type Gift struct {
	id               string
	name             string
	message          string
	receiverWalletID string
	senderWalletID   string
	baseValue        int
	status           string
	paymentID        string
	createdAt        time.Time
	updatedAt        time.Time
}

type NewGiftInput struct {
	Name             string
	Message          string
	BaseValue        int
	SenderWalletID   string
	ReceiverWalletID string
}

func NewGift(in NewGiftInput) (*Gift, error) {
	g := Gift{
		id:               core.NewID(),
		name:             in.Name,
		message:          in.Message,
		senderWalletID:   in.SenderWalletID,
		receiverWalletID: in.ReceiverWalletID,
		baseValue:        in.BaseValue,
		status:           GIFT_STATUS_PENDING,
		paymentID:        "",
		createdAt:        time.Now(),
		updatedAt:        time.Now(),
	}

	if err := g.validate(); err != nil {
		return nil, err
	}

	return &g, nil
}

type NewGiftFromInput struct {
	ID               string
	Name             string
	Message          string
	SenderWalletID   string
	ReceiverWalletID string
	BaseValue        int
	Status           string
	PaymentID        string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewGiftFrom(in NewGiftFromInput) *Gift {
	return &Gift{
		id:               in.ID,
		name:             in.Name,
		message:          in.Message,
		senderWalletID:   in.SenderWalletID,
		receiverWalletID: in.ReceiverWalletID,
		baseValue:        in.BaseValue,
		status:           in.Status,
		paymentID:        in.PaymentID,
		createdAt:        in.CreatedAt,
		updatedAt:        in.UpdatedAt,
	}
}

func (g *Gift) validate() error {
	if g.name == "" {
		return core_err.NewModelErr("name is required")
	}

	if g.senderWalletID == "" {
		return core_err.NewModelErr("senderWalletID is required")
	}

	if g.receiverWalletID == "" {
		return core_err.NewModelErr("receiverWalletID is required")
	}

	if g.baseValue < 0 {
		return core_err.NewModelErr("baseValue should be greater than 0")
	}

	return nil
}

func (g *Gift) ID() string {
	return g.id
}

func (g *Gift) Name() string {
	return g.name
}

func (g *Gift) Message() string {
	return g.message
}

func (g *Gift) BaseValue() int {
	return g.baseValue
}

func (g *Gift) Status() string {
	return g.status
}

func (g *Gift) Sent() {
	g.touch()
	g.status = GIFT_STATUS_SENT
}

func (g *Gift) SenderWalletID() string {
	return g.senderWalletID
}

func (g *Gift) ReceiverWalletID() string {
	return g.receiverWalletID
}

func (g *Gift) PaymentID() string {
	return g.paymentID
}

func (g *Gift) SetPaymentID(paymentID string) {
	g.touch()
	g.paymentID = paymentID
}

func (g *Gift) CreatedAt() time.Time {
	return g.createdAt
}

func (g *Gift) UpdatedAt() time.Time {
	return g.updatedAt
}

func (g *Gift) touch() {
	g.updatedAt = time.Now()
}
