package pages

import (
	"database/sql"
	"github.com/gchaincl/dotsql"
	"gopkg.in/macaron.v1"
	"log"
)

type portfolioItem struct {
	id          uint
	image       string
	title       string
	description string
}

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
