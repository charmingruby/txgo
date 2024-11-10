package service

import (
	"github.com/charmingruby/txgo/internal/giftshop/core/model"
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
		return err
	}

	if gift == nil {
		return core_err.NewResourceNotFoundErr("gift")
	}

	if gift.PaymentID() != "" {
		return core_err.NewResourceAlreadyExistsErr("payment")
	}

	senderWallet, err := s.walletRepo.FindByID(gift.SenderWalletID())
	if err != nil {
		return err
	}

	if senderWallet == nil {
		return core_err.NewResourceNotFoundErr("sender wallet")
	}

	partialValue := calculatePartialValue(gift.BaseValue(), params.TaxPercent, params.Installments)

	pointsDiff := senderWallet.Points() - partialValue

	if pointsDiff < 0 {
		return core_err.NewInvalidFundsErr(pointsDiff * -1)
	}

	newPaymentInput := model.NewPaymentInput{
		Installments: params.Installments,
		TaxPercent:   params.TaxPercent,
		TotalValue:   gift.BaseValue(),
	}

	payment, err := model.NewPayment(newPaymentInput)
	if err != nil {
		return err
	}

	payment.SetPartialValue(partialValue)

	if err := s.paymentRepo.Store(payment); err != nil {
		return err
	}

	transactionInput := model.NewTransactionInput{
		Points:        payment.PartialValue(),
		PayerWalletID: gift.SenderWalletID(),
	}

	transaction, err := model.NewTransaction(transactionInput)
	if err != nil {
		return err
	}

	if err := s.transactionRepo.Store(transaction); err != nil {
		return err
	}

	newWalletBalance := senderWallet.Points() - payment.PartialValue()
	senderWallet.SetPoints(newWalletBalance)
	if err := s.walletRepo.UpdatePointsByID(senderWallet); err != nil {
		return err
	}

	payment.Paid()
	payment.SetTransactionID(transaction.ID())
	if err := s.paymentRepo.UpdateTransactionIDAndStatusByID(payment); err != nil {
		return err
	}

	gift.Sent()
	gift.SetPaymentID(payment.ID())
	if err := s.giftRepo.UpdatePaymentIDAndStatusByID(gift); err != nil {
		return err
	}

	return nil
}

func calculatePartialValue(totalValue, taxPercent, installments int) int {
	totalValueWithTax := totalValue + (totalValue * taxPercent / 100)
	partialValue := totalValueWithTax / installments
	return partialValue
}
