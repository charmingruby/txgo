package response

type GiftCheckoutResponse struct {
	PaymentID     string `json:"payment_id"`
	TransactionID string `json:"transaction_id"`
}
