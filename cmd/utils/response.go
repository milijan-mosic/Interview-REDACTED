package utils

import (
	"encoding/json"
	"net/http"
)

func SetJSONResponse(message string, httpCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
}
