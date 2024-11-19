package endpoint

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/charmingruby/txgo/internal/billing/core/service"
	"github.com/charmingruby/txgo/internal/billing/transport/rest/dto/request"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/charmingruby/txgo/internal/shared/transport/rest"
	"github.com/go-chi/chi/v5"
)

func (e *Endpoint) subscribeOnPlanHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		planID := chi.URLParam(r, "planID")
		if planID == "" {
			rest.BadRequestErrorResponse[any](w, "planID path param is required")
			return
		}

		req, err := rest.ParseRequest[request.SubscribeOnPlanRequest](*e.validator, r)
		if err != nil {
			rest.BadRequestErrorResponse[any](w, err.Error())
			return
		}

		serviceInput := service.SubscribeOnPlanInput{
			Email:     req.Email,
			PlanID:    planID,
			AutoRenew: req.AutoRenew,
		}

		if err := e.service.SubscribeOnPlanService(serviceInput); err != nil {
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

			var resourceExistsErr *core_err.ResourceAlreadyExistsErr
			if errors.As(err, &resourceExistsErr) {
				rest.ConflictErrorResponse[any](w, err.Error())
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

		rest.CreatedResponse[any](w, "subscription")
	}
}
