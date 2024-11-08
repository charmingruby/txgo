package entity

import (
	"testing"

	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewPayment(t *testing.T) {
	t.Run("it should be able create a payment with valid input", func(t *testing.T) {
		input := NewPaymentInput{
			Installments: 3,
			TaxPercent:   10,
			TotalValue:   10000,
		}

		payment, err := NewPayment(input)

		assert.Nil(t, err)
		assert.NotNil(t, payment)
		assert.Equal(t, 3, payment.installments)
		assert.Equal(t, 10, payment.taxPercent)
		assert.Equal(t, 10000, payment.totalValue)
		assert.Equal(t, PAYMENT_STATUS_PENDING, payment.status)
	})

	t.Run("it should be not able create a payment with installments value less than 1", func(t *testing.T) {
		input := NewPaymentInput{
			Installments: 0,
			TaxPercent:   10,
			TotalValue:   10000,
		}

		payment, err := NewPayment(input)

		assert.Nil(t, payment)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("installments must be greater than or equal to 1").Error())
	})

	t.Run("it should be not able create a payment with tax percent value less than 0", func(t *testing.T) {
		input := NewPaymentInput{
			Installments: 1,
			TaxPercent:   -10,
			TotalValue:   10000,
		}

		payment, err := NewPayment(input)

		assert.Nil(t, payment)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("taxPercent must be greater than or equal to 0").Error())
	})

	t.Run("it should be not able create a payment with total value points less than 0", func(t *testing.T) {
		input := NewPaymentInput{
			Installments: 1,
			TaxPercent:   10,
			TotalValue:   0,
		}

		payment, err := NewPayment(input)

		assert.Nil(t, payment)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("totalValue must be greater than 0").Error())
	})
}
