package entity

import (
	"testing"

	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewPayment(t *testing.T) {
	t.Run("it should be able create a payment with valid input", func(t *testing.T) {
		input := NewPaymentInput{
			method:            PAYMENT_METHOD_CREDIT_CARD,
			installments:      3,
			taxPercent:        10,
			totalValueInCents: 10000,
			appliedPoints:     0,
		}

		payment, err := NewPayment(input)

		assert.Nil(t, err)
		assert.NotNil(t, payment)
		assert.Equal(t, PAYMENT_METHOD_CREDIT_CARD, payment.method)
		assert.Equal(t, 3, payment.installments)
		assert.Equal(t, 10, payment.taxPercent)
		assert.Equal(t, 10000, payment.totalValueInCents)
		assert.Equal(t, 0, payment.appliedPoints)
		assert.Equal(t, PAYMENT_STATUS_PENDING, payment.status)
	})

	t.Run("it should be not able create a payment with invalid payment method", func(t *testing.T) {
		input := NewPaymentInput{
			method:            "INVALID_METHOD",
			installments:      3,
			taxPercent:        10,
			totalValueInCents: 10000,
			appliedPoints:     0,
		}

		payment, err := NewPayment(input)

		assert.Nil(t, payment)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("invalid payment method").Error())
	})

	t.Run("it should be not able create a payment with installments value less than 1", func(t *testing.T) {
		input := NewPaymentInput{
			method:            PAYMENT_METHOD_CASH,
			installments:      0,
			taxPercent:        10,
			totalValueInCents: 10000,
			appliedPoints:     0,
		}

		payment, err := NewPayment(input)

		assert.Nil(t, payment)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("installments must be greater than or equal to 1").Error())
	})

	t.Run("it should be not able create a payment with tax percent value less than 0", func(t *testing.T) {
		input := NewPaymentInput{
			method:            PAYMENT_METHOD_CASH,
			installments:      1,
			taxPercent:        -10,
			totalValueInCents: 10000,
			appliedPoints:     0,
		}

		payment, err := NewPayment(input)

		assert.Nil(t, payment)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("taxPercent must be greater than or equal to 0").Error())
	})

	t.Run("it should be not able create a payment with applied points value less than 0", func(t *testing.T) {

		input := NewPaymentInput{
			method:            PAYMENT_METHOD_CASH,
			installments:      1,
			taxPercent:        10,
			totalValueInCents: 10000,
			appliedPoints:     -10,
		}

		payment, err := NewPayment(input)

		assert.Nil(t, payment)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), core_err.NewEntityErr("appliedPoints must be greater than or equal to 0").Error())
	})
}

func Test_Payment_CalculatePartialValue(t *testing.T) {
	t.Run("it should be able to calculate partial value correctly with valid input", func(t *testing.T) {
		payment := &Payment{
			method:            PAYMENT_METHOD_CREDIT_CARD,
			installments:      2,
			taxPercent:        10,
			totalValueInCents: 10000,
			appliedPoints:     0,
		}

		expectedPartialValueInCents := (payment.totalValueInCents + (payment.totalValueInCents*payment.taxPercent)/100) / payment.installments

		err := payment.CalculatePartialValue()

		assert.Nil(t, err)
		assert.Equal(t, expectedPartialValueInCents, payment.partialValueInCents)
	})

	t.Run("it should be not able to calculate partial value error if applied points are greater than total value with tax", func(t *testing.T) {
		payment := &Payment{
			method:            PAYMENT_METHOD_CREDIT_CARD,
			installments:      2,
			taxPercent:        10,
			totalValueInCents: 10000,
			appliedPoints:     200,
		}

		err := payment.CalculatePartialValue()

		assert.NotNil(t, err)
		assert.Equal(t, core_err.NewEntityErr("appliedPoints cannot be greater than totalValueInCents").Error(), err.Error())
	})

	t.Run("it should be able to calculate partial value if applied points equal total value with tax, setting partial value to 0", func(t *testing.T) {
		payment := &Payment{
			method:            PAYMENT_METHOD_CREDIT_CARD,
			installments:      1,
			taxPercent:        0,
			totalValueInCents: 10000,
			appliedPoints:     100,
		}

		err := payment.CalculatePartialValue()

		assert.Nil(t, err)
		assert.Equal(t, 0, payment.partialValueInCents)
		assert.Equal(t, PAYMENT_METHOD_POINTS, payment.method)
	})

	t.Run("it should be able to calculate partial value correctly with applied points less than total value with tax", func(t *testing.T) {
		payment := &Payment{
			method:            PAYMENT_METHOD_CREDIT_CARD,
			installments:      1,
			taxPercent:        0,
			totalValueInCents: 10000,
			appliedPoints:     50,
		}

		err := payment.CalculatePartialValue()

		assert.Nil(t, err)
		assert.Equal(t, 5000, payment.partialValueInCents)
	})

	t.Run("it should be able to calculate partial value correctly with zero tax", func(t *testing.T) {
		payment := &Payment{
			method:            PAYMENT_METHOD_CREDIT_CARD,
			installments:      2,
			taxPercent:        0,
			totalValueInCents: 10000,
			appliedPoints:     0,
		}

		err := payment.CalculatePartialValue()

		assert.Nil(t, err)
		assert.Equal(t, 5000, payment.partialValueInCents)
	})
}
