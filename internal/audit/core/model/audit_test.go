package model

import (
	"testing"

	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/stretchr/testify/assert"
)

func Test_NewAudit(t *testing.T) {
	dummyActorEmail1 := "actor1@email.com"
	dummyActorEmail2 := "actor2@email.com"
	dummyContext := "dummy context"
	dummyDescription := "dummy description"

	t.Run("it should be able to create a payment movimentation audit with valid params", func(t *testing.T) {
		actors := []string{dummyActorEmail1, dummyActorEmail2}

		input := NewAuditInput{
			actors:      actors,
			context:     dummyContext,
			description: dummyDescription,
		}

		audit, err := NewAudit(input)

		assert.Nil(t, err)
		assert.NotNil(t, audit)
		assert.Equal(t, len(actors), len(audit.actors))
		assert.Equal(t, dummyContext, audit.context)
		assert.Equal(t, dummyDescription, audit.description)
	})

	t.Run("it should not be able to create a payment movimentation audit with empty actors", func(t *testing.T) {
		input := NewAuditInput{
			actors:      []string{},
			context:     dummyContext,
			description: dummyDescription,
		}

		audit, err := NewAudit(input)

		assert.Nil(t, audit)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewModelErr("actors are required").Error(), err.Error())
	})

	t.Run("it should not be able to create a payment movimentation audit with empty context", func(t *testing.T) {
		actors := []string{dummyActorEmail1, dummyActorEmail2}

		input := NewAuditInput{
			actors:      actors,
			context:     "",
			description: dummyDescription,
		}

		audit, err := NewAudit(input)

		assert.Nil(t, audit)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewModelErr("context is required").Error(), err.Error())
	})

	t.Run("it should not be able to create a payment movimentation audit with empty description", func(t *testing.T) {
		actors := []string{dummyActorEmail1, dummyActorEmail2}

		input := NewAuditInput{
			actors:      actors,
			context:     dummyContext,
			description: "",
		}

		audit, err := NewAudit(input)

		assert.Nil(t, audit)
		assert.Error(t, err)
		assert.Equal(t, core_err.NewModelErr("description is required").Error(), err.Error())
	})
}
