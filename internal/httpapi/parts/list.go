package parts

import (
	"encoding/json"
	"interview/cmd/utils"
	"interview/internal/store"
	"log"
	"net/http"
)

func ListPartsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SetJSONResponse(w, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	rows, err := store.DB.Query("SELECT * FROM parts")
	if err != nil {
		log.Println(err)
		utils.SetJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	parts := []store.Part{}

	for rows.Next() {
		var p store.Part
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Status,
			&p.Supplier,
			&p.Material,
			&p.Weight,
			&p.Critical,
			&p.Updated,
		)
		if err != nil {
			log.Println(err)
			utils.SetJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		parts = append(parts, p)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		utils.SetJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SetJSONResponse(w, http.StatusOK, "")
	_ = json.NewEncoder(w).Encode(parts)
}
