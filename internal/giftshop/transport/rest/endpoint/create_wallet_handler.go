package endpoint

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/txgo/internal/giftshop/core/service"
	"github.com/charmingruby/txgo/internal/giftshop/transport/rest/dto/request"
	"github.com/charmingruby/txgo/internal/shared/http/rest"
)

func (e *Endpoint) createWalletHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := rest.ParseRequest[request.CreateWalletRequest](*e.validator, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		serviceInput := service.CreateWalletInput{
			WalletName:           req.WalletName,
			OwnerEmail:           req.OwnerEmail,
			InitialPointsBalance: req.InitialPointsBalance,
		}

		fmt.Printf("%+v", serviceInput)

		// if err := e.service.CreateWalletService(serviceInput); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}
