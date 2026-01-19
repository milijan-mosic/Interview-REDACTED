package parts

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"interview/cmd/utils"
	"interview/internal/store"
)

var allowedStatuses = map[string]bool{
	"Draft":     true,
	"In Review": true,
	"Approved":  true,
	"Rejected":  true,
}

type StatusUpdateRequest struct {
	Status string `json:"status"`
}

func PatchPartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		utils.SetJSONResponse(w, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") || strings.TrimPrefix(authHeader, "Bearer ") != store.AuthToken {
		utils.SetJSONResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		utils.SetJSONResponse(w, http.StatusBadRequest, "Invalid path or ID")
		return
	}
	partID := pathParts[3]

	var req StatusUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SetJSONResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	newStatus := req.Status

	if !allowedStatuses[newStatus] {
		utils.SetJSONResponse(w, http.StatusBadRequest, "Invalid status")
		return
	}

	tx, err := store.DB.Begin()
	if err != nil {
		log.Println(err)
		utils.SetJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback()

	var oldStatus string
	err = tx.QueryRowContext(r.Context(), "SELECT status FROM parts WHERE id = ?", partID).Scan(&oldStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.SetJSONResponse(w, http.StatusNotFound, err.Error())
		} else {
			utils.SetJSONResponse(w, http.StatusInternalServerError, err.Error())
		}

		log.Println(err)
		return
	}

	_, err = tx.ExecContext(r.Context(), "UPDATE parts SET status = ?, updated_at = ? WHERE id = ?", newStatus, time.Now().UTC().Format(time.RFC3339), partID)
	if err != nil {
		log.Println(err)
		utils.SetJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = tx.ExecContext(
		r.Context(),
		`INSERT INTO part_status_audit (part_id, old_status, new_status, changed_at, request_id)
		 VALUES (?, ?, ?, ?, ?)`,
		partID,
		oldStatus,
		newStatus,
		time.Now().UTC().Format(time.RFC3339),
		"",
	)
	if err != nil {
		log.Println(err)
		utils.SetJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit(); err != nil {
		log.Println(err)
		utils.SetJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SetJSONResponse(w, http.StatusOK, "")

	_ = json.NewEncoder(w).Encode(map[string]string{
		"message":    "status updated",
		"part_id":    partID,
		"old_status": oldStatus,
		"new_status": newStatus,
	})
}
