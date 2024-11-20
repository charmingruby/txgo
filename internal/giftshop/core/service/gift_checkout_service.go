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

type GiftCheckoutResult struct {
	PaymentID     string
	TransactionID string
}

func (s *Service) GiftCheckoutService(params GiftCheckoutParams) (GiftCheckoutResult, error) {
	result := GiftCheckoutResult{}

	gift, err := s.giftRepo.FindByID(params.GiftID)
	if err != nil {
		return GiftCheckoutResult{}, err
	}

	if gift == nil {
		return GiftCheckoutResult{}, core_err.NewResourceNotFoundErr("gift")
	}

	if gift.PaymentID() != "" {
		return GiftCheckoutResult{}, core_err.NewResourceAlreadyExistsErr("payment")
	}

	senderWallet, err := s.walletRepo.FindByID(gift.SenderWalletID())
	if err != nil {
		return GiftCheckoutResult{}, err
	}

	if senderWallet == nil {
		return GiftCheckoutResult{}, core_err.NewResourceNotFoundErr("sender wallet")
	}

	if !s.billingSubscriptionStatusProvider.IsSubscriptionActive(senderWallet.OwnerEmail()) {
		return GiftCheckoutResult{}, core_err.NewForbiddenActionErr("invalid user subscription")
	}

	err = s.transactionalConsistencyProvider.Transact(func(tc TransactionalConsistencyParams) error {
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
		if err := tc.PaymentRepository.Store(payment); err != nil {
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

		if err := tc.TransactionRepository.Store(transaction); err != nil {
			return err
		}

		newWalletBalance := senderWallet.Points() - payment.PartialValue()
		senderWallet.SetPoints(newWalletBalance)
		if err := tc.WalletRepository.UpdatePointsByID(senderWallet); err != nil {
			return err
		}

		payment.Paid()
		payment.SetTransactionID(transaction.ID())
		if err := tc.PaymentRepository.UpdateTransactionIDAndStatusByID(payment); err != nil {
			return err
		}

		gift.Sent()
		gift.SetPaymentID(payment.ID())
		if err := tc.GiftRepository.UpdatePaymentIDAndStatusByID(gift); err != nil {
			return err
		}

		result.PaymentID = payment.ID()
		result.TransactionID = transaction.ID()

		return nil
	})

	return result, err
}

func calculatePartialValue(totalValue, taxPercent, installments int) int {
	totalValueWithTax := totalValue + (totalValue * taxPercent / 100)
	partialValue := totalValueWithTax / installments
	return partialValue
}
