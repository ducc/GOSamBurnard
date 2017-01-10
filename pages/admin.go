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

func Admin(ctx *macaron.Context, db *sql.DB, dot *dotsql.DotSql) {
	addStandardData(ctx.Data)
	var err error
	ctx.Data["portfolio_images"], err = loadAndSortPortfolioItems(db, dot)
	ctx.Data["about_text"], ctx.Data["contact_text"] = about, contact
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx.HTMLSet(200, "base", "admin")
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
