package factory

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/charmingruby/txgo/internal/giftshop/core/model"
	"github.com/charmingruby/txgo/internal/giftshop/core/repository"
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/helper"
)

func MakeGift(giftRepo repository.GiftRepository, params model.NewGiftFromInput) (*model.Gift, error) {
	gift := createGift(params)

	if err := giftRepo.Store(gift); err != nil {
		return nil, err
	}

	return gift, nil
}

func createGift(params model.NewGiftFromInput) *model.Gift {
	input := model.NewGiftFromInput{
		ID:               helper.If[string](params.ID != "", params.ID, core.NewID()),
		Name:             helper.If[string](params.Name != "", params.Name, gofakeit.Name()),
		Message:          helper.If[string](params.Message != "", params.Message, gofakeit.Sentence(10)),
		SenderWalletID:   helper.If[string](params.SenderWalletID != "", params.SenderWalletID, "invalid-sender-wallet-id"),
		ReceiverWalletID: helper.If[string](params.ReceiverWalletID != "", params.ReceiverWalletID, "invalid-receiver-wallet-id"),
		BaseValue:        params.BaseValue,
		Status:           helper.If[string](params.Status != "", params.Status, model.GIFT_STATUS_PENDING),
		PaymentID:        helper.If[string](params.PaymentID != "", params.PaymentID, ""),
		CreatedAt:        helper.If[time.Time](params.CreatedAt != time.Time{}, params.CreatedAt, time.Now()),
		UpdatedAt:        helper.If[time.Time](params.UpdatedAt != time.Time{}, params.UpdatedAt, time.Now()),
	}

	return model.NewGiftFrom(input)
}
