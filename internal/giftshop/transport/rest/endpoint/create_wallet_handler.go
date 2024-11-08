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

func (e *Endpoint) createWalletHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := rest.ParseRequest[request.CreateWalletRequest](*e.validator, r)
		if err != nil {
			rest.BadRequestErrorResponse(w, err.Error())
			return
		}

		serviceInput := service.CreateWalletInput{
			WalletName:           req.WalletName,
			OwnerEmail:           req.OwnerEmail,
			InitialPointsBalance: req.InitialPointsBalance,
		}

		if err := e.service.CreateWalletService(serviceInput); err != nil {
			var resourceExistsErr *core_err.ResourceAlreadyExistsErr
			if errors.As(err, &resourceExistsErr) {
				rest.ConflictErrorResponse(w, err.Error())
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

		rest.CreatedResponse(w, "wallet")
	}
}
