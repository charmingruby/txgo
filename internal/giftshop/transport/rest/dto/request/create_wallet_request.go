package request

type CreateWalletRequest struct {
	WalletName           string `json:"wallet_name" validate:"required"`
	OwnerEmail           string `json:"owner_email" validate:"email"`
	InitialPointsBalance int    `json:"initial_points_balance" validate:"required"`
}
