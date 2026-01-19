package httpapi

import (
	"encoding/json"
	"interview/cmd/utils"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetJSONResponse(w, http.StatusOK, "")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
