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

func (s *Suite) Test_CreatePlanHandler() {
	route := "/plans"
	url := fmt.Sprintf("%s%s", s.server.URL, route)

	dummyPlanName := "plan name"
	dummyDescription := "plan description"
	dummyAmount := 10
	validPeriodicity := model.MONTHLY_PLAN_PERIODICITY
	dummyTrialPeriodDays := 30

	s.Run("it should be able to create a plan", func() {
		payload := request.CreatePlanRequest{
			Name:            dummyPlanName,
			Description:     dummyDescription,
			Amount:          dummyAmount,
			Periodicity:     validPeriodicity,
			TrialPeriodDays: dummyTrialPeriodDays,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url, helper.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusCreated, httpRes.StatusCode)

		decodedRes, err := helper.DecodeResponse[any](httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusCreated)
		s.Equal(decodedRes.Message, "plan created successfully")
		s.Equal(decodedRes.Data, nil)
	})

	s.Run("it should be not able to create a plan with invalid params", func() {
		payload := request.CreatePlanRequest{
			Name:            "",
			Description:     dummyDescription,
			Amount:          dummyAmount,
			Periodicity:     validPeriodicity,
			TrialPeriodDays: dummyTrialPeriodDays,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url, helper.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusBadRequest, httpRes.StatusCode)

		decodedRes, err := helper.DecodeResponse[any](httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusBadRequest)
		s.Equal(decodedRes.Message, "request validation failed: Key: 'CreatePlanRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag")
		s.Equal(decodedRes.Data, nil)
	})

	s.Run("it should be not able to create a plan with already taken name", func() {
		_, err := factory.MakePlan(s.planRepo, model.NewPlanFromInput{Name: dummyPlanName})
		s.NoError(err)

		payload := request.CreatePlanRequest{
			Name:            dummyPlanName,
			Description:     dummyDescription,
			Amount:          dummyAmount,
			Periodicity:     validPeriodicity,
			TrialPeriodDays: dummyTrialPeriodDays,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url, helper.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusConflict, httpRes.StatusCode)

		decodedRes, err := helper.DecodeResponse[any](httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusConflict)
		s.Equal(decodedRes.Message, core_err.NewResourceAlreadyExistsErr("plan").Error())
		s.Equal(decodedRes.Data, nil)
	})

	s.Run("it should be not able to create a plan with invalid peridiocity", func() {
		payload := request.CreatePlanRequest{
			Name:            dummyPlanName,
			Description:     dummyDescription,
			Amount:          dummyAmount,
			Periodicity:     "10 days",
			TrialPeriodDays: dummyTrialPeriodDays,
		}

		body, err := json.Marshal(payload)
		s.NoError(err)

		httpRes, err := http.Post(url, helper.CONTENT_TYPE_JSON, bytes.NewReader(body))
		s.NoError(err)

		s.Equal(http.StatusUnprocessableEntity, httpRes.StatusCode)

		decodedRes, err := helper.DecodeResponse[any](httpRes)
		s.NoError(err)
		s.Equal(decodedRes.Code, http.StatusUnprocessableEntity)
		s.Equal(decodedRes.Message, "invalid periodicity")
		s.Equal(decodedRes.Data, nil)
	})
}
