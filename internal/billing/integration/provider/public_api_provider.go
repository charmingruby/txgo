package provider

func NewPublicAPI() *PublicAPI {
	return &PublicAPI{}
}

type PublicAPI struct {
}

func (a *PublicAPI) IsSubscriptionActive(email string) bool {
	return false
}
