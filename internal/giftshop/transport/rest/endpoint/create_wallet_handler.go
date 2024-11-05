package endpoint

import (
	"net/http"
)

func (e *Endpoint) createWalletHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("createWalletHandler"))
	}
}
