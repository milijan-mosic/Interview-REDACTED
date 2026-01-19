package main

import (
	"interview/cmd/utils"
	"interview/internal/httpapi"
	"interview/internal/httpapi/debug"
	"interview/internal/httpapi/parts"
	"interview/internal/store"
	"log"
	"net/http"
	"os"
	"time"
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

	http.HandleFunc(store.ApiVersion+"/parts", parts.ListPartsHandler)
	http.HandleFunc(store.ApiVersion+"/parts/{id}", parts.GetPartHandler)
	http.HandleFunc(store.ApiVersion+"/parts/{id}/status", parts.PatchPartHandler)

	srv := &http.Server{
		Addr:              ":" + port,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	log.Println("Server running on :" + port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
