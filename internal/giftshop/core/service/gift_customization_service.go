package service

// type GiftCustomizationParams struct {
// 	Name          string
// 	Message       string
// 	SenderEmail   string
// 	ReceiverEmail string
// 	Value         int
// }

// func (s *Service) GiftCustomizationService(params GiftCustomizationParams) error {
// 	senderWallet, err := s.walletRepo.FindByOwnerEmail(params.SenderEmail)
// 	if err != nil {
// 		return core_err.NewResourceNotFoundErr("gift sender")
// 	}

// 	receiverWallet, err := s.walletRepo.FindByOwnerEmail(params.ReceiverEmail)
// 	if err != nil {
// 		return core_err.NewResourceNotFoundErr("gift receiver")
// 	}

// 	newGiftInput := entity.NewGiftInput{
// 		Name:          params.Name,
// 		Message:       params.Message,
// 		SenderEmail:   senderWallet.OwnerEmail(),
// 		ReceiverEmail: receiverWallet.OwnerEmail(),
// 		BaseValue:     params.Value,
// 	}

// 	newGift, err := entity.NewGift(newGiftInput)
// 	if err != nil {
// 		return err
// 	}

// 	if err := s.giftRepo.Store(newGift); err != nil {
// 		return err
// 	}

// 	return nil
// }
