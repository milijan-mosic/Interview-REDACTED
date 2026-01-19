package httpapi

import (
	"interview/cmd/utils"
	"net/http"
)

func ResetPartsHandler(w http.ResponseWriter, r *http.Request) {
	message := "Hello world!"
	utils.SetJSONResponse(w, http.StatusOK, message)
}
