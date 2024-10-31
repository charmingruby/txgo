package entity

import (
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type NewWalletInput struct {
	name       string
	ownerEmail string
}

func NewWallet(in NewWalletInput) (*Wallet, error) {
	w := Wallet{
		name:       in.name,
		ownerEmail: in.ownerEmail,
		points:     0,
		BaseEntity: core.NewBaseEntity(),
	}

	if err := w.validate(); err != nil {
		return nil, err
	}

	return &w, nil
}

func NewWalletFrom(in Wallet) *Wallet {
	return &Wallet{
		name:       in.name,
		ownerEmail: in.ownerEmail,
		points:     in.points,
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

	return nil
}

type Wallet struct {
	core.BaseEntity

	name       string
	ownerEmail string
	points     int
}
