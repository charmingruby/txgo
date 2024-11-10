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

	receiverWallet, err := s.walletRepo.FindByID(gift.ReceiverWalletID())
	if err != nil {
		return err
	}

	if receiverWallet == nil {
		return core_err.NewResourceNotFoundErr("receiver wallet")
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

	payment.SetPartialValue(calculatePartialValue(payment))

	if err := s.paymentRepo.Store(payment); err != nil {
		return err
	}

	transactionInput := model.NewTransactionInput{
		Points:           payment.PartialValue(),
		ReceiverWalletID: gift.ReceiverWalletID(),
		SenderWalletID:   gift.SenderWalletID(),
	}

	transaction, err := model.NewTransaction(transactionInput)
	if err != nil {
		return err
	}

	if err := s.transactionRepo.Store(transaction); err != nil {
		return err
	}

	receiverWallet.SetPoints(+payment.PartialValue())
	if err := s.walletRepo.UpdatePointsByID(receiverWallet); err != nil {
		return err
	}

	senderWallet.SetPoints(-payment.PartialValue())
	if err := s.walletRepo.UpdatePointsByID(senderWallet); err != nil {
		return err
	}

	payment.Paid()
	payment.SetTransactionID(transaction.ID())
	if err := s.paymentRepo.UpdateTransactionIDAndStatusByID(payment); err != nil {
		return err
	}

	gift.SetPaymentID(payment.ID())
	gift.Sent()
	if err := s.giftRepo.UpdatePaymentIDAndStatusByID(gift); err != nil {
		return err
	}

	return nil
}

func calculatePartialValue(payment *model.Payment) int {
	totalValueWithTax := payment.TotalValue() + (payment.TotalValue() * payment.TaxPercent() / 100)
	partialValue := totalValueWithTax / payment.Installments()
	return partialValue
}
