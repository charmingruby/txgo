package endpoint

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/charmingruby/txgo/internal/giftshop/core/service"
	"github.com/charmingruby/txgo/internal/giftshop/transport/rest/dto/request"
	"github.com/charmingruby/txgo/internal/giftshop/transport/rest/dto/response"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/charmingruby/txgo/internal/shared/transport/rest"
	"github.com/go-chi/chi/v5"
)

func (e *Endpoint) giftCheckoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		giftID := chi.URLParam(r, "giftID")
		if giftID == "" {
			rest.BadRequestErrorResponse[any](w, "giftID path param is required")
			return
		}

		req, err := rest.ParseRequest[request.GiftCheckoutRequest](*e.validator, r)
		if err != nil {
			rest.BadRequestErrorResponse[any](w, err.Error())
			return
		}

		serviceInput := service.GiftCheckoutParams{
			GiftID:       giftID,
			TaxPercent:   req.TaxPercent,
			Installments: req.Installments,
		}

		result, err := e.service.GiftCheckoutService(serviceInput)
		if err != nil {
			var validationErr *core_err.ModelErr
			if errors.As(err, &validationErr) {
				rest.ModelValidationErrorResponse[any](w, err.Error())
				return
			}

			var notFoundErr *core_err.ResourceNotFoundErr
			if errors.As(err, &notFoundErr) {
				rest.NotFoundErrorResponse[any](w, err.Error())
				return
			}

			var alreadyExistsErr *core_err.ResourceAlreadyExistsErr
			if errors.As(err, &alreadyExistsErr) {
				rest.ConflictErrorResponse[any](w, alreadyExistsErr.Error())
				return
			}

			var invalidFundsErr *core_err.InvalidFundsErr
			if errors.As(err, &invalidFundsErr) {
				rest.ForbiddenErrorResponse[any](w, invalidFundsErr.Error())
				return
			}

			var forbiddenActionErr *core_err.ForbiddenActionErr
			if errors.As(err, &forbiddenActionErr) {
				rest.ForbiddenErrorResponse[any](w, forbiddenActionErr.Error())
				return
			}

			var storageErr *core_err.PersistenceErr
			if errors.As(err, &storageErr) {
				slog.Error(fmt.Sprintf("PERSISTENCE ERROR: %s", storageErr.Error()))
				rest.InternalServerErrorResponse[any](w)
				return
			}

			slog.Error(fmt.Sprintf("UNEXPECTED ERROR: %s", err.Error()))
			rest.InternalServerErrorResponse[any](w)
			return
		}

		data := response.GiftCheckoutResponse{
			PaymentID:     result.PaymentID,
			TransactionID: result.TransactionID,
		}

		rest.OkResponse[response.GiftCheckoutResponse](w, "gift checkout success", data)
	}
}
