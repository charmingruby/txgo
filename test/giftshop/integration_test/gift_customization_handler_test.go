package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmingruby/txgo/internal/giftshop/core/model"
	"github.com/charmingruby/txgo/internal/giftshop/transport/rest/dto/request"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/charmingruby/txgo/test/giftshop/factory"
	"github.com/charmingruby/txgo/test/shared/helper"
)

func (s *Suite) Test_GiftCustomizationHandler() {
	route := "/gifts/customize"
	url := fmt.Sprintf("%s%s", s.server.URL, route)

	s.Run("it should be able to create a gift customization", func() {
		receiverWallet, err := factory.MakeWallet(s.walletRepo, model.NewWalletFromInput{})
		s.NoError(err)

		senderWallet, err := factory.MakeWallet(s.walletRepo, model.NewWalletFromInput{})
		s.NoError(err)

		payload := request.GiftCustomizationRequest{
			Name:          "Birthday Gift",
			Message:       "Happy birthday!",
			SenderEmail:   senderWallet.OwnerEmail(),
			ReceiverEmail: receiverWallet.OwnerEmail(),
			Value:         10000,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url, helper.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusCreated, httpRes.StatusCode)

		decodedRes, err := helper.DecodeResponse[any](httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusCreated)
		s.Equal(decodedRes.Message, "gift request created successfully")
		s.Equal(decodedRes.Data, nil)
	})

	s.Run("it should be not able to create a gift with invalid payload", func() {
		payload := request.GiftCustomizationRequest{
			Name:          "Birthday Gift",
			Message:       "Happy birthday!",
			SenderEmail:   "",
			ReceiverEmail: "dummy@mail.com",
			Value:         10000,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url, helper.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusBadRequest, httpRes.StatusCode)

		decodedRes, err := helper.DecodeResponse[any](httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusBadRequest)
		s.Equal(decodedRes.Message, "request validation failed: Key: 'GiftCustomizationRequest.SenderEmail' Error:Field validation for 'SenderEmail' failed on the 'email' tag")
		s.Equal(decodedRes.Data, nil)
	})

	s.Run("it should be not able to create a gift with an invalid owner email for sender wallet", func() {
		receiverWallet, err := factory.MakeWallet(s.walletRepo, model.NewWalletFromInput{})
		s.NoError(err)

		payload := request.GiftCustomizationRequest{
			Name:          "Birthday Gift",
			Message:       "Happy birthday!",
			SenderEmail:   "dummy@mail.com",
			ReceiverEmail: receiverWallet.OwnerEmail(),
			Value:         10000,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url, helper.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusNotFound, httpRes.StatusCode)

		decodedRes, err := helper.DecodeResponse[any](httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusNotFound)
		s.Equal(decodedRes.Message, core_err.NewResourceNotFoundErr("gift sender").Error())
		s.Equal(decodedRes.Data, nil)
	})

	s.Run("it should be not able to create a gift with an invalid owner email for receiver wallet", func() {
		senderWallet, err := factory.MakeWallet(s.walletRepo, model.NewWalletFromInput{})
		s.NoError(err)

		payload := request.GiftCustomizationRequest{
			Name:          "Birthday Gift",
			Message:       "Happy birthday!",
			ReceiverEmail: "dummy@mail.com",
			SenderEmail:   senderWallet.OwnerEmail(),
			Value:         10000,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url, helper.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusNotFound, httpRes.StatusCode)

		decodedRes, err := helper.DecodeResponse[any](httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusNotFound)
		s.Equal(decodedRes.Message, core_err.NewResourceNotFoundErr("gift receiver").Error())
		s.Equal(decodedRes.Data, nil)
	})
}
