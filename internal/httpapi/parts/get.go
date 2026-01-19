package parts

import (
	"database/sql"
	"encoding/json"
	"interview/cmd/utils"
	"interview/internal/store"
	"log"
	"net/http"
	"strings"
)

func GetPartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SetJSONResponse(w, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	id := strings.TrimPrefix(r.URL.Path, store.ApiVersion+"/parts/")
	if id == "" {
		utils.SetJSONResponse(w, http.StatusBadRequest, "Please provide part ID")
		return
	}

	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") || strings.TrimPrefix(authHeader, "Bearer ") != store.AuthToken {
		utils.SetJSONResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}

	part := store.Part{}
	query := "SELECT * FROM parts WHERE id = ?"

	err := store.DB.QueryRowContext(r.Context(), query, id).Scan(
		&part.ID,
		&part.Name,
		&part.Status,
		&part.Supplier,
		&part.Material,
		&part.Weight,
		&part.Critical,
		&part.Updated,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			utils.SetJSONResponse(w, http.StatusNotFound, "No data found")
		} else {
			log.Println(err)
			utils.SetJSONResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.SetJSONResponse(w, http.StatusOK, "")
	_ = json.NewEncoder(w).Encode(part)
}
