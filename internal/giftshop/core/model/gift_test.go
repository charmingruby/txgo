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
	dummyReceiverWallet := Wallet{
		id:         "receiver_wallet_id",
		ownerEmail: "receiver@email.com",
	}
	dummySenderWallet := Wallet{
		id:         "sender_wallet_id",
		ownerEmail: "sender@email.com",
	}

	t.Run("it should be able to create a gift with valid params", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:           dummyName,
				Message:        dummyMessage,
				SenderWallet:   &dummySenderWallet,
				ReceiverWallet: &dummyReceiverWallet,
				BaseValue:      dummyBaseValue,
			},
		)

		assert.Nil(t, err)
		assert.Equal(t, dummyName, gift.name)
		assert.Equal(t, dummyMessage, gift.message)
		assert.Equal(t, dummySenderWallet.ownerEmail, gift.SenderEmail())
		assert.Equal(t, dummyReceiverWallet.ownerEmail, gift.ReceiverEmail())
		assert.Equal(t, GIFT_STATUS_PENDING, gift.status)
	})

	t.Run("it should be not able to create a gift with empty name", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:           "",
				Message:        dummyMessage,
				SenderWallet:   &dummySenderWallet,
				ReceiverWallet: &dummyReceiverWallet,
				BaseValue:      dummyBaseValue,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewModelErr("name is required").Error())
	})

	t.Run("it should be not able to create a gift with empty senderEmail", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:           dummyName,
				Message:        dummyMessage,
				SenderWallet:   nil,
				ReceiverWallet: &dummyReceiverWallet,
				BaseValue:      dummyBaseValue,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewModelErr("senderEmail is required").Error())
	})

	t.Run("it should be not able to create a gift with empty receiverEmail", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:           dummyName,
				Message:        dummyMessage,
				SenderWallet:   &dummySenderWallet,
				ReceiverWallet: nil,
				BaseValue:      dummyBaseValue,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewModelErr("receiverEmail is required").Error())
	})

	t.Run("it should be not able to create a gift with base value less than 0", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:           dummyName,
				Message:        dummyMessage,
				SenderWallet:   &dummySenderWallet,
				ReceiverWallet: &dummyReceiverWallet,
				BaseValue:      -1,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewModelErr("baseValue should be greater than 0").Error())
	})
}
