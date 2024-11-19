package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewPlan(t *testing.T) {
	dummyName := "dummy plan name"
	dummyDescription := "dummy plan description"
	dummyAmount := 1000
	dummyInterval := "monthly"
	dummyTrialPeriodDays := 30

	t.Run("it should be able to create a plan with valid params", func(t *testing.T) {
		plan := NewPlan(
			NewPlanInput{
				Name:            dummyName,
				Description:     dummyDescription,
				Amount:          dummyAmount,
				Interval:        dummyInterval,
				TrialPeriodDays: dummyTrialPeriodDays,
			},
		)

		assert.NotNil(t, plan)
		assert.Equal(t, dummyName, plan.name)
		assert.Equal(t, dummyDescription, plan.description)
		assert.Equal(t, dummyAmount, plan.amount)
		assert.Equal(t, dummyInterval, plan.interval)
		assert.Equal(t, dummyTrialPeriodDays, plan.trialPeriodDays)
	})
}
