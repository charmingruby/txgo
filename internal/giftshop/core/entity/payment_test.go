package entity

import (
	"testing"

	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewPayment(t *testing.T) {
	t.Run("it should be able create a payment with valid input", func(t *testing.T) {
		input := NewPaymentInput{
			installments:     3,
			taxPercent:       10,
			totalValuePoints: 10000,
		}

		payment, err := NewPayment(input)

		assert.Nil(t, err)
		assert.NotNil(t, payment)
		assert.Equal(t, 3, payment.installments)
		assert.Equal(t, 10, payment.taxPercent)
		assert.Equal(t, 10000, payment.totalValuePoints)
		assert.Equal(t, PAYMENT_STATUS_PENDING, payment.status)
	})

	t.Run("it should be not able create a payment with installments value less than 1", func(t *testing.T) {
		input := NewPaymentInput{
			installments:     0,
			taxPercent:       10,
			totalValuePoints: 10000,
		}

		payment, err := NewPayment(input)

		assert.Nil(t, payment)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("installments must be greater than or equal to 1").Error())
	})

	t.Run("it should be not able create a payment with tax percent value less than 0", func(t *testing.T) {
		input := NewPaymentInput{
			installments:     1,
			taxPercent:       -10,
			totalValuePoints: 10000,
		}

		payment, err := NewPayment(input)

		assert.Nil(t, payment)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("taxPercent must be greater than or equal to 0").Error())
	})

	t.Run("it should be not able create a payment with total value points less than 0", func(t *testing.T) {
		input := NewPaymentInput{
			installments:     1,
			taxPercent:       10,
			totalValuePoints: 0,
		}

		payment, err := NewPayment(input)

		assert.Nil(t, payment)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("totalValuePoints must be greater than 0").Error())
	})
}

func Test_Payment_CalculatePartialValue(t *testing.T) {
	t.Run("it should be able to calculate partial value correctly with valid input", func(t *testing.T) {
		payment := &Payment{
			installments:     2,
			taxPercent:       10,
			totalValuePoints: 10000,
		}

		expectedPartialValueInCents := (payment.totalValuePoints + (payment.totalValuePoints*payment.taxPercent)/100) / payment.installments

		payment.CalculatePartialValue()

		assert.Equal(t, expectedPartialValueInCents, payment.partialValuePoints)
	})
}
