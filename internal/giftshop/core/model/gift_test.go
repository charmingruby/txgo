package model

import (
	"testing"

	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewGift(t *testing.T) {
	dummyName := "dummy gift name"
	dummyMessage := "dummy gift message"
	dummyBaseValue := 100
	dummyReceiverWalletID := "receiver_wallet_id"
	dummySenderWalletID := "sender_wallet_id"

	t.Run("it should be able to create a gift with valid params", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:             dummyName,
				Message:          dummyMessage,
				SenderWalletID:   dummySenderWalletID,
				ReceiverWalletID: dummyReceiverWalletID,
				BaseValue:        dummyBaseValue,
			},
		)

		assert.Nil(t, err)
		assert.Equal(t, dummyName, gift.name)
		assert.Equal(t, dummyMessage, gift.message)
		assert.Equal(t, dummySenderWalletID, gift.SenderWalletID())
		assert.Equal(t, dummyReceiverWalletID, gift.ReceiverWalletID())
		assert.Equal(t, GIFT_STATUS_PENDING, gift.status)
	})

	t.Run("it should be not able to create a gift with empty name", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:             "",
				Message:          dummyMessage,
				SenderWalletID:   dummySenderWalletID,
				ReceiverWalletID: dummyReceiverWalletID,
				BaseValue:        dummyBaseValue,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewModelErr("name is required").Error())
	})

	t.Run("it should be not able to create a gift with empty senderWalletID", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:             dummyName,
				Message:          dummyMessage,
				SenderWalletID:   "",
				ReceiverWalletID: dummyReceiverWalletID,
				BaseValue:        dummyBaseValue,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewModelErr("senderWalletID is required").Error())
	})

	t.Run("it should be not able to create a gift with empty receiverWalletID", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:             dummyName,
				Message:          dummyMessage,
				SenderWalletID:   dummySenderWalletID,
				ReceiverWalletID: "",
				BaseValue:        dummyBaseValue,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewModelErr("receiverWalletID is required").Error())
	})

	t.Run("it should be not able to create a gift with base value less than 0", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:             dummyName,
				Message:          dummyMessage,
				SenderWalletID:   dummySenderWalletID,
				ReceiverWalletID: dummyReceiverWalletID,
				BaseValue:        -1,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewModelErr("baseValue should be greater than 0").Error())
	})
}
