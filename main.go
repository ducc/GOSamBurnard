package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gchaincl/dotsql"
	"github.com/go-macaron/pongo2"
	"github.com/sponges/GOSamBurnard/pages"
	"gopkg.in/macaron.v1"
	"io/ioutil"
	"log"
	"net/http"
)

const config_path = "config.json"

type config struct {
	Database struct {
		Host     string `json:"host"`
		Database string `json:"database"`
		Username string `json:"username"`
		Password string `json:"password"`
		Ssl      bool   `json:"ssl"`
	} `json:"database"`

	Http struct {
		Port string `json:"port"`
	} `json:"http"`

	Templates struct {
		Directory string `json:"directory"`
	} `json:"templates"`
}

var (
	conf *config
	db   *sql.DB
	dot  *dotsql.DotSql
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
	dot, err = dotsql.LoadFromFile("statements.sql")
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = dot.Exec(db, "create-tables")
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
	m.Get("/", pages.Home)
	m.Get("/portfolio", pages.Portfolio)
	m.Get("/login", pages.Login)
	m.Post("/logout", pages.Logout)
	m.Get("/admin", pages.Admin)
	log.Println("Starting GOSamBurnard on", conf.Http.Port)
	err = http.ListenAndServe(conf.Http.Port, m)
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig(path string) (*config, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf config
	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

func toJSONString(input interface{}) (*string, error) {
	body, err := json.Marshal(&input)
	if err != nil {
		return nil, err
	}
	temp := string(body)
	return &temp, nil
}
