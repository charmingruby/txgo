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
)

func (e *Endpoint) createPlanHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := rest.ParseRequest[request.CreatePlanRequest](*e.validator, r)
		if err != nil {
			rest.BadRequestErrorResponse[any](w, err.Error())
			return
		}

		serviceInput := service.CreatePlanInput{
			Name:            req.Name,
			Description:     req.Description,
			Amount:          req.Amount,
			Periodicity:     req.Periodicity,
			TrialPeriodDays: req.TrialPeriodDays,
		}

		if err := e.service.CreatePlanService(serviceInput); err != nil {
			var validationErr *core_err.ModelErr
			if errors.As(err, &validationErr) {
				rest.ModelValidationErrorResponse[any](w, err.Error())
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

		rest.CreatedResponse[any](w, "plan")
	}
}
