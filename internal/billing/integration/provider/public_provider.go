package provider

func NewPublic() *Public {
	return &Public{}
}

type Public struct {
}

func (a *Public) IsSubscriptionActive(email string) bool {
	return false
}
