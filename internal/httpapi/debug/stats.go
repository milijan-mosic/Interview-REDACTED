package debug

import (
	"encoding/json"
	"interview/cmd/utils"
	"interview/internal/store"
	"log"
	"net/http"
)

func PartsStatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SetJSONResponse(w, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	partsCount := 0
	auditCount := 0

	err := store.DB.QueryRowContext(r.Context(), "SELECT COUNT(*) FROM parts").Scan(&partsCount)
	if err != nil {
		log.Println(err)
		utils.SetJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = store.DB.QueryRowContext(r.Context(), "SELECT COUNT(*) FROM part_status_audit").Scan(&auditCount)
	if err != nil {
		log.Println(err)
		utils.SetJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SetJSONResponse(w, http.StatusOK, "")

	_ = json.NewEncoder(w).Encode(map[string]int{
		"parts_count": partsCount,
		"audit_count": auditCount,
	})
}
