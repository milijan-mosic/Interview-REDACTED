package main

import (
	"fmt"
	"interview/cmd/utils"
	"interview/internal/httpapi"
	"log"
	"net/http"
	"os"
)

var dbPath string

func main() {
	http.HandleFunc("/hello", httpapi.HelloWorldHandler)

	port, err := utils.GetEnvVariable("PORT")
	if err != nil {
		log.Println(err)
		port = "8080"
	}

	dbPath, err = utils.GetEnvVariable("DATABASE_PATH")
	if err != nil {
		log.Println(err)
		dbPath = "./data.db"
	}

	log.Println(dbPath)

	log.Println("Server running on :" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
