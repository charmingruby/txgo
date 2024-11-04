package endpoint

import "net/http"

func (e *Endpoint) giftCheckoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("giftCheckoutHandler"))
	}
}
