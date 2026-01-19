package httpapi

import (
	"interview/cmd/utils"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	message := "Hello world!"
	utils.SetJSONResponse(message, http.StatusOK, w)
}
