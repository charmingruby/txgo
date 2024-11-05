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
			AmountInPoints: 100,
			ReceiverWallet: dummyWallet,
			BuyerWallet:    dummyWallet,
		}

		transaction, err := NewTransaction(input)

		assert.Nil(t, err)
		assert.NotNil(t, transaction)
		assert.Equal(t, 100, transaction.amountInPoints)
		assert.Equal(t, dummyWallet, transaction.receiverWallet)
		assert.Equal(t, dummyWallet, transaction.buyerWallet)
	})

	t.Run("it should not be able to create a transaction with amountInPoints less than or equal to 0", func(t *testing.T) {
		input := NewTransactionInput{
			AmountInPoints: 0,
			ReceiverWallet: dummyWallet,
			BuyerWallet:    dummyWallet,
		}

		transaction, err := NewTransaction(input)

		assert.Nil(t, transaction)
		assert.NotNil(t, err)
		assert.Equal(t, core_err.NewEntityErr("amountInPoints must be greater than 0").Error(), err.Error())
	})
}
