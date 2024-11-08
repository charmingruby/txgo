package entity

import (
	"time"

	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type Wallet struct {
	id         string
	name       string
	ownerEmail string
	points     int
	createdAt  time.Time
	updatedAt  time.Time
}

type NewWalletInput struct {
	Name                 string
	OwnerEmail           string
	InitialPointsBalance int
}

func NewWallet(in NewWalletInput) (*Wallet, error) {
	w := Wallet{
		id:         core.NewID(),
		name:       in.Name,
		ownerEmail: in.OwnerEmail,
		points:     in.InitialPointsBalance,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}

	if err := w.validate(); err != nil {
		return nil, err
	}

	return &w, nil
}

type NewWalletFromInput struct {
	ID         string
	Name       string
	OwnerEmail string
	Points     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewWalletFrom(in NewWalletFromInput) *Wallet {
	return &Wallet{
		id:         in.ID,
		name:       in.Name,
		ownerEmail: in.OwnerEmail,
		points:     in.Points,
		createdAt:  in.CreatedAt,
		updatedAt:  in.UpdatedAt,
	}
}

func (g *Wallet) validate() error {
	if g.name == "" {
		return core_err.NewEntityErr("name is required")
	}

	if g.ownerEmail == "" {
		return core_err.NewEntityErr("ownerEmail is required")
	}

	if g.points < 0 {
		return core_err.NewEntityErr("points must be greater than or equal to 0")
	}

	return nil
}

func (w *Wallet) ID() string {
	return w.id
}

func (w *Wallet) Name() string {
	return w.name
}

func (w *Wallet) OwnerEmail() string {
	return w.ownerEmail
}

func (w *Wallet) Points() int {
	return w.points
}

func (w *Wallet) CreatedAt() time.Time {
	return w.createdAt
}

func (w *Wallet) UpdatedAt() time.Time {
	return w.updatedAt
}
