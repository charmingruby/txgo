package integration

type BillingSubscriptionStatusAPI interface {
	IsSubscriptionActive(email string) bool
}
