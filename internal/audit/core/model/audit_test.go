package model

import (
	"testing"

	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewAudit(t *testing.T) {
	dummyModule := "dummy module"
	dummyContext := "dummy context"
	dummyMessage := "dummy message"

	t.Run("it should be able to create an audit with valid params", func(t *testing.T) {
		input := NewAuditInput{
			Module:  dummyModule,
			Context: dummyContext,
			Message: dummyMessage,
		}

		audit, err := NewAudit(input)

		assert.Nil(t, err)
		assert.NotNil(t, audit)
		assert.Equal(t, dummyModule, audit.module)
		assert.Equal(t, dummyContext, audit.context)
		assert.Equal(t, dummyMessage, audit.message)
	})

	t.Run("it should not be able to create an audit with empty context", func(t *testing.T) {
		input := NewAuditInput{
			Module:  dummyModule,
			Context: "",
			Message: dummyMessage,
		}

		audit, err := NewAudit(input)

		assert.Nil(t, audit)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewModelErr("context is required").Error(), err.Error())
	})

	t.Run("it should not be able to create an audit with empty message", func(t *testing.T) {
		input := NewAuditInput{
			Module:  dummyModule,
			Context: dummyContext,
			Message: "",
		}

		audit, err := NewAudit(input)

		assert.Nil(t, audit)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewModelErr("message is required").Error(), err.Error())
	})

	t.Run("it should not be able to create an audit with empty module", func(t *testing.T) {
		input := NewAuditInput{
			Module:  "",
			Context: dummyContext,
			Message: dummyMessage,
		}

		audit, err := NewAudit(input)

		assert.Nil(t, audit)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewModelErr("module is required").Error(), err.Error())
	})
}
