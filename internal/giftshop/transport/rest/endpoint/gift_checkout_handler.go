package endpoint

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/charmingruby/txgo/internal/giftshop/core/service"
	"github.com/charmingruby/txgo/internal/giftshop/transport/rest/dto/request"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/charmingruby/txgo/internal/shared/http/rest"
	"github.com/go-chi/chi/v5"
)

func (e *Endpoint) giftCheckoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		giftID := chi.URLParam(r, "giftID")
		if giftID == "" {
			rest.BadRequestErrorResponse(w, "giftID path param is required")
			return
		}

		req, err := rest.ParseRequest[request.GiftCheckoutRequest](*e.validator, r)
		if err != nil {
			rest.BadRequestErrorResponse(w, err.Error())
			return
		}

		serviceInput := service.GiftCheckoutParams{
			GiftID:       giftID,
			TaxPercent:   req.TaxPercent,
			Installments: req.Installments,
		}

		if err := e.service.GiftCheckoutService(serviceInput); err != nil {
			var validationErr *core_err.ModelErr
			if errors.As(err, &validationErr) {
				rest.ModelValidationErrorResponse(w, err.Error())
				return
			}

			var notFoundErr *core_err.ResourceNotFoundErr
			if errors.As(err, &notFoundErr) {
				rest.NotFoundErrorResponse(w, err.Error())
				return
			}

			var alreadyExistsErr *core_err.ResourceAlreadyExistsErr
			if errors.As(err, &alreadyExistsErr) {
				rest.ConflictErrorResponse(w, alreadyExistsErr.Error())
				return
			}

			var invalidFundsErr *core_err.InvalidFundsErr
			if errors.As(err, &invalidFundsErr) {
				rest.ForbiddenErrorResponse(w, invalidFundsErr.Error())
				return
			}

			var storageErr *core_err.PersistenceErr
			if errors.As(err, &storageErr) {
				slog.Error(fmt.Sprintf("PERSISTENCE ERROR: %s", storageErr.Error()))
				rest.InternalServerErrorResponse(w)
				return
			}

			slog.Error(fmt.Sprintf("UNEXPECTED ERROR: %s", err.Error()))
			rest.InternalServerErrorResponse(w)
			return
		}

		rest.OkResponse(w, "gift checkout success")
	}
}
