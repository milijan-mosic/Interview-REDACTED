package store

import "database/sql"

var (
	SqlFilesPathPrefix = "./cmd/sql/"
	ApiVersion         = "/v1"
	DBPath             string
	AuthToken          string
	DB                 *sql.DB
)

type Part struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Status   string  `json:"status"`
	Supplier string  `json:"supplier"`
	Material string  `json:"material"`
	Weight   float64 `json:"weight"`
	Critical bool    `json:"critical"`
	Updated  string  `json:"updated_at"`
}
