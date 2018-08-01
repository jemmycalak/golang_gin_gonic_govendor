package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToDb() (*sql.DB, error) {
	host := "127.0.0.1"
	user := "postgres"
	password := "postgre"
	dbname := "dbgolang"

	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	db, err := sql.Open("postgres", desc)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
