package main

import (
    _ "github.com/lib/pq"
    "database/sql"
    "fmt"
)

const (
    databaseDriverName = "postgres"
    databaseUriFormat = "postgres://%s:%s@%s/%s"
)

func openDatabase() (*sql.DB, error) {
    c := conf.Database
    url := fmt.Sprintf(databaseUriFormat, c.Username, c.Password, c.Host, c.Database)
    if !c.Ssl {
        url += "?sslmode=disable"
    }
    db, err := sql.Open(databaseDriverName, url)
    if err != nil {
        return nil, err
    }
    return db, nil
}