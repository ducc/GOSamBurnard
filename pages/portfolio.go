package pages

import (
	"database/sql"
	"github.com/gchaincl/dotsql"
	"gopkg.in/macaron.v1"
	"log"
	"mime/multipart"
)

type (
	AdminPortfolioEditForm struct {
		Id          int                   `form:"id"`
		Image       *multipart.FileHeader `form:"image"`
		Title       string                `form:"title"`
		Description string                `form:"description"`
		Project     int                   `form:"project"`
	}
	AdminPortfolioNewForm struct {
		Image       *multipart.FileHeader `form:"image"`
		Title       string                `form:"title"`
		Description string                `form:"description"`
		Project     int                   `form:"project"`
	}
	portfolioItem struct {
		id          uint
		image       string
		title       string
		description string
	}
)

func loadPortfolioItems(db *sql.DB, dot *dotsql.DotSql) ([]portfolioItem, error) {
	images := make([]portfolioItem, 0)
	res, err := dot.Query(db, "select-portfolio-images")
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		var (
			index   sql.NullInt64
			project sql.NullInt64
		)
		item := portfolioItem{}
		err := res.Scan(&item.id, &item.image, &item.title, &item.description, &index, &project)
		if err != nil {
			return nil, err
		}
		images = append(images, item)
	}
	return images, nil
}

func Portfolio(ctx *macaron.Context, db *sql.DB, dot *dotsql.DotSql) {
	addStandardData(ctx.Data, "portfolio")
	var err error
	ctx.Data["images"], err = loadPortfolioItems(db, dot)
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx.HTMLSet(200, "base", "portfolio")
}

func AdminPortfolioNew(ctx *macaron.Context, form AdminPortfolioNewForm, db *sql.DB, dot *dotsql.DotSql) {
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
	ctx.Redirect("/admin?alert=Created+portfolio+image!")
}

func AdminPortfolioEdit(ctx *macaron.Context, form AdminPortfolioEditForm, db *sql.DB, dot *dotsql.DotSql) {
	var (
		err  error
		name *string
	)
	if form.Image != nil {
		name, err = saveImage(form.Image)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	var x interface{}
	if form.Project != 0 {
		x = form.Project
	}
	if name != nil {
		_, err = dot.Exec(db, "update-portfolio-image", *name, form.Title, form.Description, x, form.Id)
		if err != nil {
			log.Fatalln(err)
			return
		}
	} else {
		_, err = dot.Exec(db, "update-portfolio-image-info", form.Title, form.Description, x, form.Id)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
	ctx.Redirect("/admin?alert=Edited+portfolio+image!")
}

func AdminPortfolioDelete(ctx *macaron.Context, db *sql.DB, dot *dotsql.DotSql) {
	_, err := dot.Exec(db, "delete-portfolio-image", ctx.Params("id"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	ctx.Redirect("/admin?alert=Deleted+portfolio+image!")
}
