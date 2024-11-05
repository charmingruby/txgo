package service

import (
	"github.com/charmingruby/txgo/internal/giftshop/core/entity"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

type GiftCheckoutParams struct {
	GiftID       string
	TaxPercent   int
	Installments int
}

func (s *Service) GiftCheckoutService(params GiftCheckoutParams) error {
	gift, err := s.giftRepo.FindByID(params.GiftID)
	if err != nil {
		return core_err.NewResourceNotFoundErr("gift")
	}

	receiverWallet, err := s.walletRepo.FindByOwnerEmail(gift.ReceiverEmail())
	if err != nil {
		return core_err.NewResourceNotFoundErr("receiver wallet")
	}

	buyerWallet, err := s.walletRepo.FindByOwnerEmail(gift.SenderEmail())
	if err != nil {
		return core_err.NewResourceNotFoundErr("sender wallet")
	}

	newPaymentInput := entity.NewPaymentInput{
		Installments:     params.Installments,
		TaxPercent:       params.TaxPercent,
		TotalValuePoints: gift.BaseValueInPoints(),
	}

	payment, err := entity.NewPayment(newPaymentInput)
	if err != nil {
		return err
	}

	payment.CalculatePartialValue()

	if err := s.paymentRepo.Store(payment); err != nil {
		return err
	}

	transactionInput := entity.NewTransactionInput{
		AmountInPoints: gift.BaseValueInPoints(),
		ReceiverWallet: receiverWallet,
		BuyerWallet:    buyerWallet,
	}

	transaction, err := entity.NewTransaction(transactionInput)
	if err != nil {
		return err
	}

	if err := s.transactionRepo.Store(transaction); err != nil {
		return err
	}

	return err
}
