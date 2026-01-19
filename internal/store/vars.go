package store

import "database/sql"

var (
	SqlFilesPathPrefix = "./cmd/sql/"
	DBPath             string
	AuthToken          string
	DB                 *sql.DB
)
