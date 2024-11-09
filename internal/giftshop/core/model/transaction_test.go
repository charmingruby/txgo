package model

import (
	"testing"

	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewTransaction(t *testing.T) {
	dummyReceiverWalletID := "receiver_wallet"
	dummySenderWalletID := "sender_wallet"

	t.Run("it should be able to create a transaction with valid params", func(t *testing.T) {
		input := NewTransactionInput{
			Points:           100,
			ReceiverWalletID: dummyReceiverWalletID,
			SenderWalletID:   dummySenderWalletID,
		}

		transaction, err := NewTransaction(input)

		assert.Nil(t, err)
		assert.NotNil(t, transaction)
		assert.Equal(t, 100, transaction.points)
		assert.Equal(t, dummyReceiverWalletID, transaction.receiverWalletID)
		assert.Equal(t, dummySenderWalletID, transaction.senderWalletID)
	})

	t.Run("it should not be able to create a transaction with points less than or equal to 0", func(t *testing.T) {
		input := NewTransactionInput{
			Points:           0,
			ReceiverWalletID: dummyReceiverWalletID,
			SenderWalletID:   dummySenderWalletID,
		}

		transaction, err := NewTransaction(input)

		assert.Nil(t, transaction)
		assert.NotNil(t, err)
		assert.Equal(t, core_err.NewModelErr("points must be greater than 0").Error(), err.Error())
	})
}
