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
			isInPoints:    false,
			amountInCents: 1000,
			buyerWallet:   dummyWallet,
		}

		transaction, err := NewTransaction(input)

		assert.Nil(t, err)
		assert.NotNil(t, transaction)
		assert.Equal(t, false, transaction.isInPoints)
		assert.Equal(t, 1000, transaction.amountInCents)
		assert.Equal(t, dummyWallet, transaction.buyerWallet)
	})

	t.Run("it should not be able to create a transaction with amountInCents less than or equal to 0", func(t *testing.T) {
		input := NewTransactionInput{
			isInPoints:    false,
			amountInCents: 0,
			buyerWallet:   dummyWallet,
		}

		transaction, err := NewTransaction(input)

		assert.Nil(t, transaction)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewEntityErr("amountInCents must be greater than 0").Error(), err.Error())
	})

	t.Run("it should not be able to create a transaction with nil buyerWallet", func(t *testing.T) {
		input := NewTransactionInput{
			isInPoints:    false,
			amountInCents: 1000,
			buyerWallet:   nil,
		}

		transaction, err := NewTransaction(input)

		assert.Nil(t, transaction)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewEntityErr("buyerWallet is required").Error(), err.Error())
	})
}
