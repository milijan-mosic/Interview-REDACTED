package main

import (
	"database/sql"
	"fmt"
	"interview/cmd/utils"
	"interview/internal/httpapi"
	"log"
	"net/http"
	"os"
)

var (
	dbPath    string
	authToken string
	db        *sql.DB
)

func main() {
	port := utils.InitAPI(&dbPath, &authToken)

	db, err := utils.InitDatabase(dbPath, "./cmd/sql/schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/hello", httpapi.HelloWorldHandler)

	log.Println("Server running on :" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
