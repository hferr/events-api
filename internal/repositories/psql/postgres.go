package psql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB
}

func NewPostgres(connString string) (*Postgres, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Db: db,
	}, nil
}
