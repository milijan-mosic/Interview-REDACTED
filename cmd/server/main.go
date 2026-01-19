package main

import (
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
)

func main() {
	http.HandleFunc("/hello", httpapi.HelloWorldHandler)

	port := utils.InitAPI(&dbPath, &authToken)
	log.Println(dbPath)    // debug
	log.Println(authToken) // debug

	log.Println("Server running on :" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
