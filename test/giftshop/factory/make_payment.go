package factory

import (
	"time"

	"github.com/charmingruby/txgo/internal/giftshop/core/model"
	"github.com/charmingruby/txgo/internal/giftshop/core/repository"
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/helper"
)

func MakePayment(paymentRepo repository.PaymentRepository, params model.NewPaymentFromInput) (*model.Payment, error) {
	payment := createPayment(params)

	if err := paymentRepo.Store(payment); err != nil {
		return nil, err
	}

	return payment, nil
}

func createPayment(params model.NewPaymentFromInput) *model.Payment {
	input := model.NewPaymentFromInput{
		ID:            helper.If[string](params.ID != "", params.ID, core.NewID()),
		Installments:  params.Installments,
		TaxPercent:    params.TaxPercent,
		PartialValue:  params.PartialValue,
		TotalValue:    params.TotalValue,
		Status:        helper.If[string](params.Status != "", params.Status, model.GIFT_STATUS_PENDING),
		TransactionID: helper.If[string](params.TransactionID != "", params.TransactionID, ""),
		CreatedAt:     helper.If[time.Time](params.CreatedAt != time.Time{}, params.CreatedAt, time.Now()),
		UpdatedAt:     helper.If[time.Time](params.UpdatedAt != time.Time{}, params.UpdatedAt, time.Now()),
	}

	return model.NewPaymentFrom(input)
}
