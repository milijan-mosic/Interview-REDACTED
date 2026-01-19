package debug

import (
	"interview/cmd/utils"
	"interview/internal/store"
	"log"
	"net/http"
	"os"
)

func ResetPartsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.SetJSONResponse(w, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	sqlFiles := []string{"clean_db.sql", "schema.sql", "seed.sql"}

	for _, file := range sqlFiles {
		schema, err := os.ReadFile(store.SqlFilesPathPrefix + file)
		if err != nil {
			log.Println(err)
			utils.SetJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		_, err = utils.ExecSQL(store.DB, schema)
		if err != nil {
			log.Println(err)
			utils.SetJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	utils.SetJSONResponse(w, http.StatusOK, "Database is reset")
}
