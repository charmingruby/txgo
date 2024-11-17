package model

import (
	"testing"

	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewAuditActor(t *testing.T) {
	dummyEmail := "dummy@example.com"
	dummyAuditID := "dummyAuditID"

	t.Run("it should be able to create an audit actor with valid params", func(t *testing.T) {
		input := NewAuditActorInput{
			Email:   dummyEmail,
			AuditID: dummyAuditID,
		}

		auditActor, err := NewAuditActor(input)

		assert.Nil(t, err)
		assert.NotNil(t, auditActor)
		assert.Equal(t, dummyEmail, auditActor.email)
		assert.Equal(t, dummyAuditID, auditActor.auditID)
	})

	t.Run("it should not be able to create an audit actor with empty email", func(t *testing.T) {
		input := NewAuditActorInput{
			Email:   "",
			AuditID: dummyAuditID,
		}

		auditActor, err := NewAuditActor(input)

		assert.Nil(t, auditActor)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewModelErr("email is required").Error(), err.Error())
	})

	t.Run("it should not be able to create an audit actor with empty auditID", func(t *testing.T) {
		input := NewAuditActorInput{
			Email:   dummyEmail,
			AuditID: "",
		}

		auditActor, err := NewAuditActor(input)

		assert.Nil(t, auditActor)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewModelErr("auditID is required").Error(), err.Error())
	})
}
