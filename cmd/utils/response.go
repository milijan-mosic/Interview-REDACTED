package utils

import (
	"net/http"
)

func SetJSONResponse(w http.ResponseWriter, httpCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
}
