package giftshop_integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmingruby/txgo/internal/giftshop/core/model"
	"github.com/charmingruby/txgo/internal/giftshop/transport/rest/dto/request"
	"github.com/charmingruby/txgo/internal/giftshop/transport/rest/dto/response"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/charmingruby/txgo/test/factory"
	"github.com/charmingruby/txgo/test/integration"
)

func (s *Suite) Test_GiftCheckoutHandler_Base() {
	url := func(id string) string {
		return fmt.Sprintf("%s/gifts/%s/checkout", s.server.URL, id)
	}

	s.Run("it should be able to checkout a payment", func() {
		tax := 10
		giftValue := 1000
		giftValueWithTax := giftValue + (giftValue * tax / 100)

		walletExtraBalance := 1
		walletBaseBalance := giftValueWithTax + walletExtraBalance

		senderWallet, err := factory.MakeWallet(s.walletRepo, model.NewWalletFromInput{
			Points: walletBaseBalance,
		})
		s.NoError(err)

		receiverWallet, err := factory.MakeWallet(s.walletRepo, model.NewWalletFromInput{})
		s.NoError(err)

		gift, err := factory.MakeGift(s.giftRepo, model.NewGiftFromInput{
			SenderWalletID:   senderWallet.ID(),
			ReceiverWalletID: receiverWallet.ID(),
			BaseValue:        giftValue,
		})
		s.NoError(err)

		payload := request.GiftCheckoutRequest{
			TaxPercent:   tax,
			Installments: 1,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url(gift.ID()), integration.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusOK, httpRes.StatusCode)

		decodedRes, err := integration.DecodeResponse[response.GiftCheckoutResponse](httpRes)

		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusOK)
		s.Equal(decodedRes.Message, "gift checkout success")

		payment, err := s.paymentRepo.FindByID(decodedRes.Data.PaymentID)
		s.NoError(err)
		s.Equal(model.PAYMENT_STATUS_PAID, payment.Status())
		s.Equal(giftValueWithTax, payment.PartialValue())
		s.Equal(decodedRes.Data.TransactionID, payment.TransactionID())

		transaction, err := s.transactionRepo.FindByID(decodedRes.Data.TransactionID)
		s.NoError(err)
		s.Equal(giftValueWithTax, transaction.Points())

		modifiedWallet, err := s.walletRepo.FindByID(senderWallet.ID())
		modifiedWalletBalance := walletBaseBalance - giftValueWithTax
		s.NoError(err)
		s.Equal(modifiedWalletBalance, modifiedWallet.Points())

		modifiedGift, err := s.giftRepo.FindByID(gift.ID())
		s.NoError(err)
		s.Equal(model.GIFT_STATUS_SENT, modifiedGift.Status())
		s.Equal(decodedRes.Data.PaymentID, modifiedGift.PaymentID())
	})

	s.Run("it should be not able to checkout a payment with invalid payload", func() {
		senderWallet, err := factory.MakeWallet(s.walletRepo, model.NewWalletFromInput{
			Points: 100,
		})
		s.NoError(err)

		receiverWallet, err := factory.MakeWallet(s.walletRepo, model.NewWalletFromInput{})
		s.NoError(err)

		gift, err := factory.MakeGift(s.giftRepo, model.NewGiftFromInput{
			SenderWalletID:   senderWallet.ID(),
			ReceiverWalletID: receiverWallet.ID(),
			BaseValue:        100,
		})
		s.NoError(err)

		payload := request.GiftCheckoutRequest{
			TaxPercent:   -10,
			Installments: 1,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url(gift.ID()), integration.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusBadRequest, httpRes.StatusCode)

		decodedRes, err := integration.DecodeResponse[response.GiftCheckoutResponse](httpRes)

		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusBadRequest)
		s.Equal(decodedRes.Message, "request validation failed: Key: 'GiftCheckoutRequest.TaxPercent' Error:Field validation for 'TaxPercent' failed on the 'min' tag")
	})

	s.Run("it should be not able to checkout a gift if doesn't exists", func() {
		payload := request.GiftCheckoutRequest{
			TaxPercent:   10,
			Installments: 1,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url("ïnvalid-id"), integration.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusNotFound, httpRes.StatusCode)

		decodedRes, err := integration.DecodeResponse[response.GiftCheckoutResponse](httpRes)

		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusNotFound)
		s.Equal(decodedRes.Message, core_err.NewResourceNotFoundErr("gift").Error())
	})

	s.Run("it should be not able to checkout a gift if already have a payment attached", func() {
		tax := 10
		giftValue := 1000
		giftValueWithTax := giftValue + (giftValue * tax / 100)

		walletExtraBalance := 1
		walletBaseBalance := giftValueWithTax + walletExtraBalance

		senderWallet, err := factory.MakeWallet(s.walletRepo, model.NewWalletFromInput{
			Points: walletBaseBalance,
		})
		s.NoError(err)

		receiverWallet, err := factory.MakeWallet(s.walletRepo, model.NewWalletFromInput{})
		s.NoError(err)

		payment, err := factory.MakePayment(s.paymentRepo, model.NewPaymentFromInput{})
		s.NoError(err)

		gift, err := factory.MakeGift(s.giftRepo, model.NewGiftFromInput{
			SenderWalletID:   senderWallet.ID(),
			ReceiverWalletID: receiverWallet.ID(),
			BaseValue:        giftValue,
			PaymentID:        payment.ID(),
		})
		s.NoError(err)

		payload := request.GiftCheckoutRequest{
			TaxPercent:   tax,
			Installments: 1,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url(gift.ID()), integration.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusConflict, httpRes.StatusCode)

		decodedRes, err := integration.DecodeResponse[response.GiftCheckoutResponse](httpRes)

		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusConflict)
		s.Equal(decodedRes.Message, core_err.NewResourceAlreadyExistsErr("payment").Error())
	})

	s.Run("it should be not able to checkout a gift if wallet doesn't have enough points", func() {
		giftValue := 1000
		walletBaseBalance := giftValue - 1

		senderWallet, err := factory.MakeWallet(s.walletRepo, model.NewWalletFromInput{
			Points: walletBaseBalance,
		})
		s.NoError(err)

		receiverWallet, err := factory.MakeWallet(s.walletRepo, model.NewWalletFromInput{})
		s.NoError(err)

		gift, err := factory.MakeGift(s.giftRepo, model.NewGiftFromInput{
			SenderWalletID:   senderWallet.ID(),
			ReceiverWalletID: receiverWallet.ID(),
			BaseValue:        giftValue,
		})
		s.NoError(err)

		payload := request.GiftCheckoutRequest{
			TaxPercent:   0,
			Installments: 1,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url(gift.ID()), integration.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusForbidden, httpRes.StatusCode)

		decodedRes, err := integration.DecodeResponse[response.GiftCheckoutResponse](httpRes)

		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusForbidden)
		s.Equal(decodedRes.Message, core_err.NewInvalidFundsErr(1).Error())
	})

}

func (s *Suite) Test_GiftCheckoutHandler_Transaction() {}
