package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmingruby/txgo/internal/billing/core/model"
	"github.com/charmingruby/txgo/internal/billing/transport/rest/dto/request"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/charmingruby/txgo/test/billing/factory"
	"github.com/charmingruby/txgo/test/shared/helper"
)

func (s *Suite) Test_SubscribeOnPlanHandler() {
	url := func(id string) string {
		return fmt.Sprintf("%s/plans/%s/subscribe", s.server.URL, id)
	}

	dummyEmail := "dummy@email.com"

	s.Run("it should be able to subscribe on a plan", func() {
		plan, err := factory.MakePlan(s.planRepo, model.NewPlanFromInput{})
		s.NoError(err)

		payload := request.SubscribeOnPlanRequest{
			Email:     dummyEmail,
			AutoRenew: false,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url(plan.ID()), helper.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusCreated, httpRes.StatusCode)

		decodedRes, err := helper.DecodeResponse[any](httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusCreated)
		s.Equal(decodedRes.Message, "subscription created successfully")
		s.Equal(decodedRes.Data, nil)
	})

	s.Run("it should be not able to subscribe on a plan with invalid payload", func() {
		plan, err := factory.MakePlan(s.planRepo, model.NewPlanFromInput{})
		s.NoError(err)

		payload := request.SubscribeOnPlanRequest{
			Email:     "",
			AutoRenew: false,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url(plan.ID()), helper.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusBadRequest, httpRes.StatusCode)

		decodedRes, err := helper.DecodeResponse[any](httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusBadRequest)
		s.Equal(decodedRes.Message, "request validation failed: Key: 'SubscribeOnPlanRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag")
		s.Equal(decodedRes.Data, nil)
	})

	s.Run("it should be not able to subscribe on a plan with invalid plan_id", func() {
		payload := request.SubscribeOnPlanRequest{
			Email:     dummyEmail,
			AutoRenew: false,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url("invalid-id"), helper.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusNotFound, httpRes.StatusCode)

		decodedRes, err := helper.DecodeResponse[any](httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusNotFound)
		s.Equal(decodedRes.Message, core_err.NewResourceNotFoundErr("plan").Error())
		s.Equal(decodedRes.Data, nil)
	})

	s.Run("it should be not able to subscribe on a plan already subscribed", func() {
		plan, err := factory.MakePlan(s.planRepo, model.NewPlanFromInput{})
		s.NoError(err)

		_, err = factory.MakeSubscription(s.subscriptionRepo, model.NewSubscriptionFromInput{
			PlanID: plan.ID(),
			Email:  dummyEmail,
		})
		s.NoError(err)

		payload := request.SubscribeOnPlanRequest{
			Email:     dummyEmail,
			AutoRenew: false,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url(plan.ID()), helper.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusConflict, httpRes.StatusCode)

		decodedRes, err := helper.DecodeResponse[any](httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusConflict)
		s.Equal(decodedRes.Message, core_err.NewResourceAlreadyExistsErr("subscription").Error())
		s.Equal(decodedRes.Data, nil)
	})
}
