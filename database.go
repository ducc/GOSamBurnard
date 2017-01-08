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

func sslModeString(ssl bool) string {
	if ssl {
		return "enable"
	} else {
		return "disable"
	}
}

func openDatabase(conf *config) (*sql.DB, error) {
	c := conf.Database
	url := fmt.Sprintf(database_uri_format, c.Username, c.Password, c.Host, c.Port, c.Database, sslModeString(c.Ssl))
	db, err := sql.Open(database_driver_name, url)
	if err != nil {
		return nil, err
	}
	return db, nil
}
