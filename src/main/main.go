package main

import (
    "log"
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/pongo2"
    "net/http"
    "database/sql"
)

const config_path = "config.json"

var (
    conf *config
    db *sql.DB
)

func main() {
    var err error
    conf, err = loadConfig(config_path)
    if err != nil {
        log.Fatal(err)
        return
    }
    db, err = openDatabase()
    if err != nil {
        log.Fatal(err)
        return
    }
    m := macaron.Classic()
    m.Use(macaron.Static("static", macaron.StaticOptions{
        Prefix: "static",
    }))
    m.Use(pongo2.Pongoers(pongo2.Options{
        Directory: conf.Templates.Directory,
    }, "base:templates"))
    m.Get("/", home)
    m.Get("/portfolio", portfolio)
    err = http.ListenAndServe(conf.Http.Port, m)
    if err != nil {
        log.Fatal(err)
    }
}