package entity

import (
	"testing"

	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewTransaction(t *testing.T) {
	dummyWallet := &Wallet{
		name:       "dummy wallet",
		ownerEmail: "owner@email.com",
		points:     1000,
		BaseEntity: core.NewBaseEntity(),
	}

	t.Run("it should be able to create a transaction with valid params", func(t *testing.T) {
		input := NewTransactionInput{
			Points:         100,
			ReceiverWallet: dummyWallet,
			SenderWallet:   dummyWallet,
		}

		transaction, err := NewTransaction(input)

		assert.Nil(t, err)
		assert.NotNil(t, transaction)
		assert.Equal(t, 100, transaction.points)
		assert.Equal(t, dummyWallet, transaction.receiverWallet)
		assert.Equal(t, dummyWallet, transaction.senderWallet)
	})

	t.Run("it should not be able to create a transaction with points less than or equal to 0", func(t *testing.T) {
		input := NewTransactionInput{
			Points:         0,
			ReceiverWallet: dummyWallet,
			SenderWallet:   dummyWallet,
		}

		transaction, err := NewTransaction(input)

		assert.Nil(t, transaction)
		assert.NotNil(t, err)
		assert.Equal(t, core_err.NewEntityErr("points must be greater than 0").Error(), err.Error())
	})
}
