package pages

import (
	"database/sql"
	"github.com/gchaincl/dotsql"
	"gopkg.in/macaron.v1"
	"log"
	"mime/multipart"
	"sort"
	"strconv"
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
		index       int
	}
	portfolioItems []portfolioItem
)

func (items portfolioItems) Len() int {
	return len(items)
}

func (items portfolioItems) Less(i, j int) bool {
	return items[i].index < items[j].index
}

func (items portfolioItems) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func loadPortfolioItems(db *sql.DB, dot *dotsql.DotSql) (portfolioItems, error) {
	images := make(portfolioItems, 0)
	res, err := dot.Query(db, "select-portfolio-images")
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		var project sql.NullInt64
		item := portfolioItem{}
		err := res.Scan(&item.id, &item.image, &item.title, &item.description, &item.index, &project)
		if err != nil {
			return nil, err
		}
		images = append(images, item)
	}
	return images, nil
}

func loadAndSortPortfolioItems(db *sql.DB, dot *dotsql.DotSql) (portfolioItems, error) {
	images, err := loadPortfolioItems(db, dot)
	if err != nil {
		return nil, err
	}
	sort.Sort(images)
	return images, nil
}

func Portfolio(ctx *macaron.Context, db *sql.DB, dot *dotsql.DotSql) {
	addStandardData(ctx.Data, "portfolio")
	var err error
	ctx.Data["images"], err = loadAndSortPortfolioItems(db, dot)
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
	res, err := dot.Query(db, "select-portfolio-images-max-index")
	var index int64
	for res.Next() {
		var nullIndex sql.NullInt64
		err = res.Scan(&nullIndex)
		if err != nil {
			log.Fatalln(err)
			return
		}
		if !nullIndex.Valid {
			index = 0
		} else {
			i, err := nullIndex.Value()
			if err != nil {
				log.Fatalln(err)
				return
			}
			index = i.(int64)
		}
	}
	var x interface{}
	if form.Project != 0 {
		x = form.Project
	}
	_, err = dot.Exec(db, "insert-portfolio-image", name, form.Title, form.Description, index+1, x)
	if err != nil {
		log.Fatalln(err)
		return
	}
	ctx.Redirect("/admin?alert=Created+portfolio+image!#admin-portfolio")
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
	ctx.Redirect("/admin?alert=Edited+portfolio+image!#admin-portfolio")
}

func AdminPortfolioDelete(ctx *macaron.Context, db *sql.DB, dot *dotsql.DotSql) {
	_, err := dot.Exec(db, "delete-portfolio-image", ctx.Params("id"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	ctx.Redirect("/admin?alert=Deleted+portfolio+image!#admin-portfolio")
}

func AdminPortfolioOrder(ctx *macaron.Context, db *sql.DB, dot *dotsql.DotSql) {
	action := ctx.Params(":action")
	index, err := strconv.Atoi(ctx.Params(":index"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	switch action {
	case "up":
		index = index - 1
		break
	case "down":
		index = index + 1
		break
	}
	_, err = dot.Exec(db, "update-portfolio-image-order", index, ctx.Params("id"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	ctx.Redirect("/admin#admin-portfolio")
}
