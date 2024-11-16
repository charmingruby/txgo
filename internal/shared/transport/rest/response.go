package rest

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func newResponse[T any](w http.ResponseWriter, code int, message string, data T) {
	w.Header().Add("Content-Type", "application/json")

	res := Response[T]{
		Code:    code,
		Message: message,
		Data:    data,
	}

	jsonRes, _ := json.Marshal(res)

	w.WriteHeader(code)
	w.Write(jsonRes)
}

func BadRequestErrorResponse[T any](w http.ResponseWriter, message string) {
	newResponse[*T](w, http.StatusBadRequest, message, nil)
}

func NotFoundErrorResponse[T any](w http.ResponseWriter, message string) {
	newResponse[*T](w, http.StatusNotFound, message, nil)
}

func ModelValidationErrorResponse[T any](w http.ResponseWriter, message string) {
	newResponse[*T](w, http.StatusUnprocessableEntity, message, nil)
}

func ConflictErrorResponse[T any](w http.ResponseWriter, message string) {
	newResponse[*T](w, http.StatusConflict, message, nil)
}

func ForbiddenErrorResponse[T any](w http.ResponseWriter, message string) {
	newResponse[*T](w, http.StatusForbidden, message, nil)
}

func InternalServerErrorResponse[T any](w http.ResponseWriter) {
	newResponse[*T](w, http.StatusInternalServerError, "internal server error", nil)
}

func CreatedResponse[T any](w http.ResponseWriter, createdResource string) {
	newResponse[*T](w, http.StatusCreated, createdResource+" created successfully", nil)
}

func OkResponse[T any](w http.ResponseWriter, message string, data T) {
	newResponse[*T](w, http.StatusOK, message, &data)
}
