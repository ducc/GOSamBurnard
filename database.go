package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	database_driver_name = "postgres"
	database_uri_format  = "postgres://%s:%s@%s:%s/%s?sslmode=%s"
)

func sslModeString() string {
	if conf.Database.Ssl {
		return "enable"
	} else {
		return "disable"
	}
}

func openDatabase() (*sql.DB, error) {
	c := conf.Database
	url := fmt.Sprintf(database_uri_format, c.Username, c.Password, c.Host, c.Port, c.Database, sslModeString())
	db, err := sql.Open(database_driver_name, url)
	if err != nil {
		return nil, err
	}
	return db, nil
}
