package utils

import (
	"encoding/json"
	"net/http"
)

func SetJSONResponse(w http.ResponseWriter, httpCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	if message != "" {
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": message,
		})
	}
}
