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
)

func (e *Endpoint) giftCustomizationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := rest.ParseRequest[request.GiftCustomizationRequest](*e.validator, r)
		if err != nil {
			rest.BadRequestErrorResponse(w, err.Error())
			return
		}

		serviceInput := service.GiftCustomizationParams{
			Name:          req.Name,
			Message:       req.Message,
			SenderEmail:   req.SenderEmail,
			ReceiverEmail: req.ReceiverEmail,
			Value:         req.Value,
		}

		if err := e.service.GiftCustomizationService(serviceInput); err != nil {
			var notFoundErr *core_err.ResourceNotFoundErr
			if errors.As(err, &notFoundErr) {
				rest.NotFoundErrorResponse(w, err.Error())
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

		rest.CreatedResponse(w, "gift request")
	}
}
