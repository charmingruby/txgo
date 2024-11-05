package entity

import (
	"testing"

	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewWallet(t *testing.T) {
	dummyName := "dummy wallet name"
	dummyOwnerEmail := "owner@email.com"
	dummyPoints := 1000

	t.Run("it should be able to create a wallet with valid params", func(t *testing.T) {
		wallet, err := NewWallet(
			NewWalletInput{
				Name:                 dummyName,
				OwnerEmail:           dummyOwnerEmail,
				InitialPointsBalance: dummyPoints,
			},
		)

		assert.Nil(t, err)
		assert.NotNil(t, wallet)
		assert.Equal(t, dummyName, wallet.name)
		assert.Equal(t, dummyOwnerEmail, wallet.ownerEmail)
		assert.Equal(t, dummyPoints, wallet.points)
	})

	t.Run("it should not be able to create a wallet with empty name", func(t *testing.T) {
		wallet, err := NewWallet(
			NewWalletInput{
				Name:                 "",
				OwnerEmail:           dummyOwnerEmail,
				InitialPointsBalance: dummyPoints,
			},
		)

		assert.Nil(t, wallet)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewEntityErr("name is required").Error(), err.Error())
	})

	t.Run("it should not be able to create a wallet with empty ownerEmail", func(t *testing.T) {
		wallet, err := NewWallet(
			NewWalletInput{
				Name:                 dummyName,
				OwnerEmail:           "",
				InitialPointsBalance: dummyPoints,
			},
		)

		assert.Nil(t, wallet)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewEntityErr("ownerEmail is required").Error(), err.Error())
	})

	t.Run("it should not be able to create a wallet with negative points", func(t *testing.T) {
		wallet, err := NewWallet(
			NewWalletInput{
				Name:                 dummyName,
				OwnerEmail:           dummyOwnerEmail,
				InitialPointsBalance: -10,
			},
		)

		assert.Nil(t, wallet)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewEntityErr("points must be greater than or equal to 0").Error(), err.Error())
	})
}
