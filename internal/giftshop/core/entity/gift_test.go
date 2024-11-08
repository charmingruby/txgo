package entity

import (
	"testing"

	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewGift(t *testing.T) {
	dummyName := "dummy gift name"
	dummyMessage := "dummy gift message"
	dummyReceiverEmail := "receiver@email.com"
	dummySenderEmail := "sender@email.com"
	dummyBaseValue := 100

	t.Run("it should be able to create a gift with valid params", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:          dummyName,
				Message:       dummyMessage,
				SenderEmail:   dummySenderEmail,
				ReceiverEmail: dummyReceiverEmail,
				BaseValue:     dummyBaseValue,
			},
		)

		assert.Nil(t, err)
		assert.Equal(t, dummyName, gift.name)
		assert.Equal(t, dummyMessage, gift.message)
		assert.Equal(t, dummySenderEmail, gift.senderEmail)
		assert.Equal(t, dummyReceiverEmail, gift.receiverEmail)
		assert.Equal(t, GIFT_STATUS_PENDING, gift.status)
	})

	t.Run("it should be not able to create a gift with empty name", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:          "",
				Message:       dummyMessage,
				SenderEmail:   dummySenderEmail,
				ReceiverEmail: dummyReceiverEmail,
				BaseValue:     dummyBaseValue,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("name is required").Error())
	})

	t.Run("it should be not able to create a gift with empty senderEmail", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:          dummyName,
				Message:       dummyMessage,
				SenderEmail:   "",
				ReceiverEmail: dummyReceiverEmail,
				BaseValue:     dummyBaseValue,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("senderEmail is required").Error())
	})

	t.Run("it should be not able to create a gift with empty receiverEmail", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:          dummyName,
				Message:       dummyMessage,
				SenderEmail:   dummySenderEmail,
				ReceiverEmail: "",
				BaseValue:     dummyBaseValue,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("receiverEmail is required").Error())
	})

	t.Run("it should be not able to create a gift with base value less than 0", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				Name:          dummyName,
				Message:       dummyMessage,
				SenderEmail:   dummySenderEmail,
				ReceiverEmail: dummyReceiverEmail,
				BaseValue:     -1,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("baseValue should be greater than 0").Error())
	})
}
