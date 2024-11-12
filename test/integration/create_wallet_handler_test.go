package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmingruby/txgo/internal/giftshop/transport/rest/dto/request"
)

func (s *Suite) Test_CreateWalletHandler() {
	route := "/wallets"
	url := fmt.Sprintf("%s%s", s.server.URL, route)

	s.Run("it should be able to create a new wallet", func() {
		payload := request.CreateWalletRequest{
			WalletName:           "My Wallet",
			OwnerEmail:           "my_wallet@email.com",
			InitialPointsBalance: 1000,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url, CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusCreated, httpRes.StatusCode)

		decodedRes, err := decodeResponse(httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusCreated)
		s.Equal(decodedRes.Message, "wallet created successfully")
		s.Equal(decodedRes.Data, nil)

		wallet, err := s.walletRepo.FindByOwnerEmail(payload.OwnerEmail)
		s.NoError(err)
		s.Equal(payload.WalletName, wallet.Name())
		s.Equal(payload.OwnerEmail, wallet.OwnerEmail())
		s.Equal(payload.InitialPointsBalance, wallet.Points())
	})
}
