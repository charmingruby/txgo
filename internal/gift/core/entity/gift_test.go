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

	t.Run("it should be able to create a gift with valid params", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				name:          dummyName,
				message:       dummyMessage,
				senderEmail:   dummySenderEmail,
				receiverEmail: dummyReceiverEmail,
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
				name:          "",
				message:       dummyMessage,
				senderEmail:   dummySenderEmail,
				receiverEmail: dummyReceiverEmail,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("name is required").Error())
	})

	t.Run("it should be not able to create a gift with empty senderEmail", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				name:          dummyName,
				message:       dummyMessage,
				senderEmail:   "",
				receiverEmail: dummyReceiverEmail,
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("senderEmail is required").Error())
	})

	t.Run("it should be not able to create a gift with empty receiverEmail", func(t *testing.T) {
		gift, err := NewGift(
			NewGiftInput{
				name:          dummyName,
				message:       dummyMessage,
				senderEmail:   dummySenderEmail,
				receiverEmail: "",
			},
		)

		assert.Nil(t, gift)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("receiverEmail is required").Error())
	})
}
