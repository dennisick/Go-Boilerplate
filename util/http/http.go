package http

import (
	"encoding/json"
	"net/http"
)

type HttpErrorResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

func Error(w http.ResponseWriter, error string, code int) {
	json.NewEncoder(w).Encode(HttpErrorResponse{
		Code:   code,
		Status: http.StatusText(code),
		Error:  error,
	})
}

func BadRequest(w http.ResponseWriter, error string) {
	Error(w, error, http.StatusBadRequest)
}

func Unauthorized(w http.ResponseWriter, error string) {
	Error(w, error, http.StatusUnauthorized)
}

func Forbidden(w http.ResponseWriter, error string) {
	Error(w, error, http.StatusForbidden)
}

func NotFound(w http.ResponseWriter, error string) {
	Error(w, error, http.StatusNotFound)
}

func InternalServerError(w http.ResponseWriter, error string) {
	Error(w, error, http.StatusInternalServerError)
}
