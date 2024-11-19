package helper

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/charmingruby/txgo/internal/shared/transport/rest"
)

const (
	CONTENT_TYPE_JSON = "application/json"
)

func DecodeResponse[T any](r *http.Response) (rest.Response[T], error) {
	resBody := r.Body
	defer resBody.Close()

	body, err := io.ReadAll(resBody)
	if err != nil {
		return rest.Response[T]{}, err
	}

	var response rest.Response[T]
	err = json.Unmarshal(body, &response)
	return response, err
}
