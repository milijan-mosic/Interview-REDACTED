package parts

import (
	"interview/cmd/utils"
	"net/http"
	"strings"
)

func GetPartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.SetJSONResponse(w, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/v1/parts/")
	if id == "" {
		utils.SetJSONResponse(w, http.StatusBadRequest, "Please provide part ID")
		return
	}

	utils.SetJSONResponse(w, http.StatusOK, string(id))
}
