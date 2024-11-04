package endpoint

import "net/http"

func (e *Endpoint) giftCustomizationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("giftCustomizationHandler"))
	}
}
