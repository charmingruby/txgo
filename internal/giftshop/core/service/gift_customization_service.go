package service

import (
	"github.com/charmingruby/txgo/internal/giftshop/core/entity"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type GiftCustomizationParams struct {
	Name          string
	Message       string
	SenderEmail   string
	ReceiverEmail string
	ValueInPoints int
}

func (s *Service) GiftCustomizationService(params GiftCustomizationParams) error {
	senderWallet, err := s.walletRepo.FindByOwnerEmail(params.SenderEmail)
	if err != nil {
		return core_err.NewResourceNotFoundErr("gift sender")
	}

	receiverWallet, err := s.walletRepo.FindByOwnerEmail(params.ReceiverEmail)
	if err != nil {
		return core_err.NewResourceNotFoundErr("gift receiver")
	}

	newGiftInput := entity.NewGiftInput{
		Name:              params.Name,
		Message:           params.Message,
		SenderEmail:       senderWallet.OwnerEmail(),
		ReceiverEmail:     receiverWallet.OwnerEmail(),
		BaseValueInPoints: params.ValueInPoints,
	}

	newGift, err := entity.NewGift(newGiftInput)
	if err != nil {
		return err
	}

	if err := s.giftRepo.Store(newGift); err != nil {
		return err
	}

	return nil
}
