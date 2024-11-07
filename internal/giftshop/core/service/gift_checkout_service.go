package service

// type GiftCheckoutParams struct {
// 	GiftID       string
// 	TaxPercent   int
// 	Installments int
// }

// func (s *Service) GiftCheckoutService(params GiftCheckoutParams) error {
// 	gift, err := s.giftRepo.FindByID(params.GiftID)
// 	if err != nil {
// 		return core_err.NewResourceNotFoundErr("gift")
// 	}

// 	receiverWallet, err := s.walletRepo.FindByOwnerEmail(gift.ReceiverEmail())
// 	if err != nil {
// 		return core_err.NewResourceNotFoundErr("receiver wallet")
// 	}

// 	senderWallet, err := s.walletRepo.FindByOwnerEmail(gift.SenderEmail())
// 	if err != nil {
// 		return core_err.NewResourceNotFoundErr("sender wallet")
// 	}

// 	newPaymentInput := entity.NewPaymentInput{
// 		Installments: params.Installments,
// 		TaxPercent:   params.TaxPercent,
// 		TotalValue:   gift.BaseValue(),
// 	}

// 	payment, err := entity.NewPayment(newPaymentInput)
// 	if err != nil {
// 		return err
// 	}

// 	payment.CalculatePartialValue()

// 	if err := s.paymentRepo.Store(payment); err != nil {
// 		return err
// 	}

// 	transactionInput := entity.NewTransactionInput{
// 		Points:         gift.BaseValue(),
// 		ReceiverWallet: receiverWallet,
// 		SenderWallet:   senderWallet,
// 	}

// 	transaction, err := entity.NewTransaction(transactionInput)
// 	if err != nil {
// 		return err
// 	}

// 	if err := s.transactionRepo.Store(transaction); err != nil {
// 		return err
// 	}

// 	return err
// }
