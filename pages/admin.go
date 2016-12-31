package pages

import (
	"database/sql"
	"github.com/gchaincl/dotsql"
	"gopkg.in/macaron.v1"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
)

const image_file_path = "static/img/"

type (
	AdminPortfolioForm struct {
		Image       *multipart.FileHeader `form:"image"`
		Title       string                `form:"title"`
		Description string                `form:"description"`
		Project     int                   `form:"project"`
	}
)

func Admin(ctx *macaron.Context, db *sql.DB, dot *dotsql.DotSql) {
	addStandardData(ctx.Data)
	var err error
	ctx.Data["portfolio_images"], err = loadPortfolioItems(db, dot)
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx.HTMLSet(200, "base", "admin")
}

func AdminPortfolioNew(ctx *macaron.Context, form AdminPortfolioForm, db *sql.DB, dot *dotsql.DotSql) {
	name, err := saveImage(form.Image)
	if err != nil {
		log.Fatal(err)
		return
	}
	var x interface{}
	if form.Project != 0 {
		x = form.Project
	}
	_, err = dot.Exec(db, "insert-portfolio-image", name, form.Title, form.Description, x)
	if err != nil {
		log.Fatalln(err)
		return
	}
	ctx.Redirect("/admin")
}

func saveImage(header *multipart.FileHeader) (*string, error) {
	name := image_file_path + header.Filename
	file, err := header.Open()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(name, body, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	return &name, nil
}
