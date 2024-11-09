package service

import (
	"github.com/charmingruby/txgo/internal/giftshop/core/model"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type GiftCustomizationParams struct {
	Name          string
	Message       string
	SenderEmail   string
	ReceiverEmail string
	Value         int
}

func (s *Service) GiftCustomizationService(params GiftCustomizationParams) error {
	senderWallet, err := s.walletRepo.FindByOwnerEmail(params.SenderEmail)
	if err != nil {
		return err
	}

	if senderWallet == nil {
		return core_err.NewResourceNotFoundErr("gift sender")
	}

	receiverWallet, err := s.walletRepo.FindByOwnerEmail(params.ReceiverEmail)
	if err != nil {
		return err
	}

	if receiverWallet == nil {
		return core_err.NewResourceNotFoundErr("gift receiver")
	}

	newGiftInput := model.NewGiftInput{
		Name:             params.Name,
		Message:          params.Message,
		SenderWalletID:   senderWallet.ID(),
		ReceiverWalletID: receiverWallet.ID(),
		BaseValue:        params.Value,
	}

	newGift, err := model.NewGift(newGiftInput)
	if err != nil {
		return err
	}

	if err := s.giftRepo.Store(newGift); err != nil {
		return err
	}

	return nil
}
