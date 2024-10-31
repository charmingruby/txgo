package entity

import (
	"testing"

	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewPaymentMovimentationAudit(t *testing.T) {
	dummyActorEmail := "actor@email.com"
	dummyContext := "dummy context"
	dummyAmountInCents := 1000

	t.Run("it should be able to create a payment movimentation audit with valid params", func(t *testing.T) {
		input := NewPaymentMovimentationAuditInput{
			actorEmail:    dummyActorEmail,
			context:       dummyContext,
			amountInCents: dummyAmountInCents,
		}

		audit, err := NewPaymentMovimentationAudit(input)

		assert.Nil(t, err)
		assert.NotNil(t, audit)
		assert.Equal(t, dummyActorEmail, audit.actorEmail)
		assert.Equal(t, dummyContext, audit.context)
		assert.Equal(t, dummyAmountInCents, audit.amountInCents)
	})

	t.Run("it should not be able to create a payment movimentation audit with empty actorEmail", func(t *testing.T) {
		input := NewPaymentMovimentationAuditInput{
			actorEmail:    "",
			context:       dummyContext,
			amountInCents: dummyAmountInCents,
		}

		audit, err := NewPaymentMovimentationAudit(input)

		assert.Nil(t, audit)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewEntityErr("actorEmail is required").Error(), err.Error())
	})

	t.Run("it should not be able to create a payment movimentation audit with empty context", func(t *testing.T) {
		input := NewPaymentMovimentationAuditInput{
			actorEmail:    dummyActorEmail,
			context:       "",
			amountInCents: dummyAmountInCents,
		}

		audit, err := NewPaymentMovimentationAudit(input)

		assert.Nil(t, audit)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewEntityErr("context is required").Error(), err.Error())
	})

	t.Run("it should not be able to create a payment movimentation audit with amountInCents less than or equal to 0", func(t *testing.T) {
		input := NewPaymentMovimentationAuditInput{
			actorEmail:    dummyActorEmail,
			context:       dummyContext,
			amountInCents: 0,
		}

		audit, err := NewPaymentMovimentationAudit(input)

		assert.Nil(t, audit)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewEntityErr("amountInCents must be greater than 0").Error(), err.Error())
	})
}
