package request

type GiftCheckoutRequest struct {
	TaxPercent   int `json:"tax_percent" validate:"min=0"`
	Installments int `json:"installments" validate:"min=1"`
}
