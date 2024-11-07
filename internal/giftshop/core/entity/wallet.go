package entity

import (
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type NewWalletInput struct {
	Name                 string
	OwnerEmail           string
	InitialPointsBalance int
}

func NewWallet(in NewWalletInput) (*Wallet, error) {
	w := Wallet{
		name:       in.Name,
		ownerEmail: in.OwnerEmail,
		points:     in.InitialPointsBalance,
		BaseEntity: core.NewBaseEntity(),
	}

	if err := w.validate(); err != nil {
		return nil, err
	}

	return &w, nil
}

type NewWalletFromInput struct {
	BaseEntity core.BaseEntity
	Name       string
	OwnerEmail string
	Points     int
}

func NewWalletFrom(in NewWalletFromInput) *Wallet {
	return &Wallet{
		name:       in.Name,
		ownerEmail: in.OwnerEmail,
		points:     in.Points,
		BaseEntity: in.BaseEntity,
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

func (w *Wallet) OwnerEmail() string {
	return w.ownerEmail
}

type Wallet struct {
	core.BaseEntity

	name       string
	ownerEmail string
	points     int
}
