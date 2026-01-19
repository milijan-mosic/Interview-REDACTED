package parts

import (
	"interview/cmd/utils"
	"net/http"
)

func PatchPartHandler(w http.ResponseWriter, r *http.Request) {
	message := "Hello world!"
	utils.SetJSONResponse(w, http.StatusOK, message)
}
