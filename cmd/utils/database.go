package utils

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func ExecSQL(db *sql.DB, schema []byte) (any, error) {
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

	return db, nil
}
