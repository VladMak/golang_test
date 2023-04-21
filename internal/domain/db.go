package domain

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DB interface {
	Connect() (*sql.DB, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	CreateDb(username, password, host, dbname string, port int)
}
