package store

import "database/sql"

var (
	DBPath    string
	AuthToken string
	DB        *sql.DB
)
