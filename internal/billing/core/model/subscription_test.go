package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewSubscription(t *testing.T) {
	dummyEmail := "dummy@example.com"
	dummyPlanID := "dummyPlanID"
	dummyStartDate := time.Now()
	dummyEndDate := time.Now().AddDate(0, 1, 0)
	dummyEndDatePtr := &dummyEndDate
	dummyAutoRenew := true

	t.Run("it should be able to create a subscription with valid params", func(t *testing.T) {
		input := NewSubscriptionInput{
			Email:     dummyEmail,
			PlanID:    dummyPlanID,
			StartDate: dummyStartDate,
			EndDate:   dummyEndDatePtr,
			AutoRenew: dummyAutoRenew,
		}

		subscription := NewSubscription(input)

		assert.NotNil(t, subscription)
		assert.Equal(t, dummyEmail, subscription.email)
		assert.Equal(t, dummyPlanID, subscription.planID)
		assert.Equal(t, SUBSCRIPTION_STATUS_PENDING, subscription.status)
		assert.Equal(t, dummyStartDate, subscription.startDate)
		assert.Equal(t, dummyEndDatePtr, subscription.endDate)
		assert.Equal(t, dummyAutoRenew, subscription.autoRenew)
	})
}
