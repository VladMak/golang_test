package domain

import (
	"database/sql"
	"fmt"
)

type PostgresDB struct {
	*sql.DB
	username string
	password string
	host     string
	port     int
	dbname   string
}

func (pg *PostgresDB) CreateDb(username, password, host, dbname string, port int) {
	pg.username = username
	pg.password = password
	pg.host =  host
	pg.port = port
	pg.dbname = dbname
}

func (pg *PostgresDB) Connect() (*sql.DB, error) {
	// Connect to the postgres db using the connection string.
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		pg.username, pg.password, pg.host, pg.port, pg.dbname))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (pg *PostgresDB) Exec(path string) {
	// Execute a query with the given args.
	pg.DB.Exec(fmt.Sprintf("insert into changes (path) values ('%s')", path))
}

func (pg *PostgresDB) Close() error {
	// Close the connection to the postgres db.
	return pg.DB.Close()
}
