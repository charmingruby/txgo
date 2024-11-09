package request

type GiftCustomizationRequest struct {
	Name          string `json:"name" validate:"required"`
	Message       string `json:"message"`
	SenderEmail   string `json:"sender_email" validate:"email"`
	ReceiverEmail string `json:"receiver_email" validate:"email"`
	Value         int    `json:"value" validate:"min=0"`
}
