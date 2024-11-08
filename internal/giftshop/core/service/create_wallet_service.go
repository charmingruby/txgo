package service

import (
	"github.com/charmingruby/txgo/internal/giftshop/core/entity"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type CreateWalletInput struct {
	WalletName           string
	OwnerEmail           string
	InitialPointsBalance int
}

func (s *Service) CreateWalletService(params CreateWalletInput) error {
	walletFound, err := s.walletRepo.FindByOwnerEmail(params.OwnerEmail)
	if err != nil {
		return err
	}

	if walletFound != nil {
		return core_err.NewResourceAlreadyExistsErr("wallet")
	}

	input := entity.NewWalletInput{
		Name:                 params.WalletName,
		OwnerEmail:           params.OwnerEmail,
		InitialPointsBalance: params.InitialPointsBalance,
	}

	newWallet, err := entity.NewWallet(input)
	if err != nil {
		return err
	}

	if err := s.walletRepo.Store(newWallet); err != nil {
		return err
	}

	return nil
}
