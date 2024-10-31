package entity

import (
	"testing"

	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewWallet(t *testing.T) {
	dummyName := "dummy wallet name"
	dummyOwnerEmail := "owner@email.com"

	t.Run("it should be able to create a wallet with valid params", func(t *testing.T) {
		wallet, err := NewWallet(
			NewWalletInput{
				name:       dummyName,
				ownerEmail: dummyOwnerEmail,
			},
		)

		assert.Nil(t, err)
		assert.NotNil(t, wallet)
		assert.Equal(t, dummyName, wallet.name)
		assert.Equal(t, dummyOwnerEmail, wallet.ownerEmail)
		assert.Equal(t, 0, wallet.points)
	})

	t.Run("it should not be able to create a wallet with empty name", func(t *testing.T) {
		wallet, err := NewWallet(
			NewWalletInput{
				name:       "",
				ownerEmail: dummyOwnerEmail,
			},
		)

		assert.Nil(t, wallet)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewEntityErr("name is required").Error(), err.Error())
	})

	t.Run("it should not be able to create a wallet with empty ownerEmail", func(t *testing.T) {
		wallet, err := NewWallet(
			NewWalletInput{
				name:       dummyName,
				ownerEmail: "",
			},
		)

		assert.Nil(t, wallet)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewEntityErr("ownerEmail is required").Error(), err.Error())
	})
}
