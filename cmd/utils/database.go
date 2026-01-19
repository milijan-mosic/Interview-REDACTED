package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func ExecSQL(db *sql.DB, schema []byte) (sql.Result, error) {
	if result, err := db.Exec(string(schema)); err != nil {
		return result, err
	}

	return nil, nil
}

func InitDatabase(dbPath string, schemaPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(string(schema)); err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping DB:", err)
	}

	return db, nil
}
