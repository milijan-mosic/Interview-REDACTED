package main

import (
	"fmt"
	"interview/cmd/utils"
	"interview/internal/httpapi"
	"interview/internal/httpapi/debug"
	"interview/internal/store"
	"log"
	"net/http"
	"os"
)

func main() {
	port := utils.InitAPI(&store.DBPath, &store.AuthToken)

	db, err := utils.InitDatabase(store.DBPath, store.SqlFilesPathPrefix+"schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	store.DB = db
	defer store.DB.Close()

	seedSchema, err := os.ReadFile(store.SqlFilesPathPrefix + "seed.sql")
	if err != nil {
		log.Fatal(err)
	}
	_, err = utils.ExecSQL(db, seedSchema)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Seeding database finished...")
	}

	http.HandleFunc("/health", httpapi.HealthHandler)

	http.HandleFunc("/debug/reset", debug.ResetPartsHandler)
	http.HandleFunc("/debug/stats", debug.PartsStatsHandler)

	log.Println("Server running on :" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
