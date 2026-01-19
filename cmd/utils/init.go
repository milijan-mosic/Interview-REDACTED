package utils

import (
	"log"
)

func InitAPI(dbPath *string, authToken *string) string {
	port, err := GetEnvVariable("PORT")
	if err != nil {
		log.Println(err)
		port = "8080"
	}
	*dbPath, err = GetEnvVariable("DATABASE_PATH")
	if err != nil {
		log.Println(err)
		*dbPath = "./data.db"
	}
	*authToken, err = GetEnvVariable("AUTH_TOKEN")
	if err != nil {
		log.Println(err)
		*authToken = "dev-token"
	}

	return port
}
