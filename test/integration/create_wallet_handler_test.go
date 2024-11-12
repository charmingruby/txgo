package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmingruby/txgo/internal/giftshop/core/model"
	"github.com/charmingruby/txgo/internal/giftshop/transport/rest/dto/request"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/charmingruby/txgo/test/factory"
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

	s.Run("it should be not able to create a new wallet with an invalid payload", func() {
		payload := request.CreateWalletRequest{
			WalletName:           "My Wallet",
			OwnerEmail:           "",
			InitialPointsBalance: 1000,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url, CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusBadRequest, httpRes.StatusCode)

		decodedRes, err := decodeResponse(httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusBadRequest)
		s.Equal(decodedRes.Message, "request validation failed: Key: 'CreateWalletRequest.OwnerEmail' Error:Field validation for 'OwnerEmail' failed on the 'email' tag")
		s.Equal(decodedRes.Data, nil)
	})

	s.Run("it should be not able to create a new wallet with an existing owner email", func() {
		conflictingEmail := "my_wallet@email.com"

		_, err := factory.MakeWallet(s.walletRepo, factory.MakeWalletParams{
			Input: model.NewWalletFromInput{
				OwnerEmail: conflictingEmail,
			},
		})
		s.NoError(err)

		payload := request.CreateWalletRequest{
			WalletName:           "My Wallet",
			OwnerEmail:           conflictingEmail,
			InitialPointsBalance: 1000,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url, CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusConflict, httpRes.StatusCode)

		decodedRes, err := decodeResponse(httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusConflict)
		s.Equal(decodedRes.Message, core_err.NewResourceAlreadyExistsErr("wallet").Error())
		s.Equal(decodedRes.Data, nil)
	})

	s.Run("it should be not able to create a new wallet with negative initial points", func() {
		payload := request.CreateWalletRequest{
			WalletName:           "My Wallet",
			OwnerEmail:           "my_wallet@email.com",
			InitialPointsBalance: -1,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url, CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusUnprocessableEntity, httpRes.StatusCode)

		decodedRes, err := decodeResponse(httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusUnprocessableEntity)
		s.Equal(decodedRes.Message, "points must be greater than or equal to 0")
		s.Equal(decodedRes.Data, nil)
	})
}
