package rest

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func newResponse(w http.ResponseWriter, code int, message string, data any) {
	w.Header().Add("Content-Type", "application/json")

	res := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	jsonRes, _ := json.Marshal(res)

	w.WriteHeader(code)
	w.Write(jsonRes)
}

func BadRequestErrorResponse(w http.ResponseWriter, message string) {
	newResponse(w, http.StatusBadRequest, message, nil)
}

func NotFoundErrorResponse(w http.ResponseWriter, message string) {
	newResponse(w, http.StatusNotFound, message, nil)
}

func ModelValidationErrorResponse(w http.ResponseWriter, message string) {
	newResponse(w, http.StatusUnprocessableEntity, message, nil)
}

func ConflictErrorResponse(w http.ResponseWriter, message string) {
	newResponse(w, http.StatusConflict, message, nil)
}

func ForbiddenErrorResponse(w http.ResponseWriter, message string) {
	newResponse(w, http.StatusForbidden, message, nil)
}

func InternalServerErrorResponse(w http.ResponseWriter) {
	newResponse(w, http.StatusInternalServerError, "internal server error", nil)
}

func CreatedResponse(w http.ResponseWriter, createdResource string) {
	newResponse(w, http.StatusCreated, createdResource+" created succesfully", nil)
}

func OkResponse(w http.ResponseWriter, message string) {
	newResponse(w, http.StatusOK, message, nil)
}
