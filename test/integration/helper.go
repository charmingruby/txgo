package integration

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/charmingruby/txgo/internal/shared/http/rest"
)

const (
	CONTENT_TYPE_JSON = "application/json"
)

func decodeResponse(r *http.Response) (rest.Response, error) {
	resBody := r.Body
	defer resBody.Close()

	body, err := io.ReadAll(resBody)
	if err != nil {
		return rest.Response{}, err
	}

	var response rest.Response
	err = json.Unmarshal(body, &response)
	return response, err
}
