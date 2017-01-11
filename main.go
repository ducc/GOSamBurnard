package main

import (
	"encoding/json"
	"fmt"
	"github.com/gchaincl/dotsql"
	"github.com/go-macaron/binding"
	"github.com/go-macaron/pongo2"
	"github.com/go-macaron/session"
	_ "github.com/go-macaron/session/postgres"
	"github.com/sponges/GOSamBurnard/pages"
	"gopkg.in/macaron.v1"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	config_path             = "config.json"
	session_provider_format = "user=%s password=%s host=%s port=%s dbname=%s sslmode=%s"
)

type (
	config struct {
		Database struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
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
)

func main() {
	conf, err := loadConfig(config_path)
	if err != nil {
		log.Fatal(err)
		return
	}
	db, err := openDatabase(conf)
	if err != nil {
		log.Fatal(err)
		return
	}
	dot, err := dotsql.LoadFromFile("statements.sql")
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
	m.Map(db)
	m.Map(dot)
	m.Use(session.Sessioner(session.Options{
		Provider: database_driver_name,
		ProviderConfig: fmt.Sprintf(session_provider_format, conf.Database.Username, conf.Database.Password,
			conf.Database.Host, conf.Database.Port, conf.Database.Database, sslModeString(conf.Database.Ssl)),
	}))
	m.Use(macaron.Static("static", macaron.StaticOptions{
		Prefix: "static",
	}))
	m.Use(pongo2.Pongoers(pongo2.Options{
		Directory: conf.Templates.Directory,
	}, "base:templates"))
	m.Use(func(ctx *macaron.Context) {
		if len(ctx.QueryStrings("alert")) != 0 {
			ctx.Data["show_alert"] = true
			ctx.Data["alert"] = ctx.QueryStrings("alert")[0]
		}
	})
	pages.Init()
	m.Get("/", pages.Home)
	m.Get("/portfolio", pages.Portfolio)
	m.Get("/projects", pages.Projects)
	//m.Get("/projects/:id", pages.Project)
	m.Get("/about", pages.Information)
	m.Get("/login", pages.Login)
	m.Post("/logout", pages.Logout)
	m.Get("/admin", pages.Admin)
	m.Post("/admin/portfolio/new", binding.MultipartForm(pages.AdminPortfolioNewForm{}), pages.AdminPortfolioNew)
	m.Post("/admin/portfolio/edit", binding.MultipartForm(pages.AdminPortfolioEditForm{}), pages.AdminPortfolioEdit)
	m.Get("/admin/portfolio/delete/:id", pages.AdminPortfolioDelete)
	m.Get("/admin/portfolio/order/:id/:index/:action", pages.AdminPortfolioOrder)
	m.Post("/admin/information", binding.Form(pages.InformationForm{}), pages.AdminInformation)
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
