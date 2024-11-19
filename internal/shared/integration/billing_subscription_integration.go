package integration

type BillingSubscriptionStatusIntegration interface {
	IsSubscriptionActive(email string) bool
}
