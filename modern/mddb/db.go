package mddb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgresDb(host string, port int, user string, password string, dbName string) (*sql.DB, error) {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
