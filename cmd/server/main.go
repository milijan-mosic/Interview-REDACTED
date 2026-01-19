package main

import (
	"fmt"
	"interview/cmd/utils"
	"interview/internal/httpapi"
	"interview/internal/store"
	"log"
	"net/http"
	"os"
)

func main() {
	port := utils.InitAPI(&store.DBPath, &store.AuthToken)

	db, err := utils.InitDatabase(store.DBPath, "./cmd/sql/schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	seedSchema, err := os.ReadFile("./cmd/sql/seed.sql")
	if err != nil {
		log.Fatal(err)
	}
	err = utils.ExecSQL(db, seedSchema)
	if err != nil {
		log.Println(err)
	}

	http.HandleFunc("/health", httpapi.HealthHandler)

	log.Println("Server running on :" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
