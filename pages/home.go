package pages

import (
	"gopkg.in/macaron.v1"
	"mime/multipart"
	"github.com/gchaincl/dotsql"
	"database/sql"
	"log"
	"sort"
	"strconv"
)

type (
	AdminSliderNewForm struct {
		Image *multipart.FileHeader `form:"image"`
	}

	AdminSliderEditForm struct {
		Id string `form:"id"`
		Image *multipart.FileHeader `form:"image"`
	}

	sliderItem struct {
		id uint
		image string
		index int
	}

	sliderItems []sliderItem
)

func (items sliderItems) Len() int {
	return len(items)
}

func (items sliderItems) Less(i, j int) bool {
	return items[i].index < items[j].index
}

func (items sliderItems) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func loadSliderItems(db *sql.DB, dot *dotsql.DotSql) (sliderItems, error) {
	items := make(sliderItems, 0)
	res, err := dot.Query(db, "select-home-images")
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		item := sliderItem{}
		err = res.Scan(&item.id, &item.image, &item.index)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func loadAndSortSliderItems(db *sql.DB, dot *dotsql.DotSql) (sliderItems, error) {
	items, err := loadSliderItems(db, dot)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	sort.Sort(items)
	return items, nil
}

func Home(ctx *macaron.Context, db *sql.DB, dot *dotsql.DotSql) {
	addStandardData(ctx.Data, "home")
	var err error
	ctx.Data["slider_items"], err = loadAndSortSliderItems(db, dot)
	if err != nil {
		log.Fatalln(err)
		return
	}
	ctx.HTMLSet(200, "base", "index")
}

func AdminSliderNew(ctx *macaron.Context, form AdminSliderNewForm, db *sql.DB, dot *dotsql.DotSql) {
	imagePath, err := saveImage(form.Image)
	if err != nil {
		log.Fatal(err)
		return
	}
	res, err := dot.Query(db, "select-home-images-max-index")
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
	_, err = dot.Exec(db, "insert-home-image", imagePath, index+1)
	if err != nil {
		log.Fatalln(err)
		return
	}
	ctx.Redirect("/admin?alert=Added+slider+image!#admin-homepage")
}

func AdminSliderEdit(ctx *macaron.Context, form AdminSliderEditForm, db *sql.DB, dot *dotsql.DotSql) {
	imagePath, err := saveImage(form.Image)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = dot.Exec(db, "update-home-image", imagePath, form.Id)
	if err != nil {
		log.Fatalln(err)
		return
	}
	ctx.Redirect("/admin?alert=Updated+slider+image!#admin-homepage")
}

func AdminSliderDelete(ctx *macaron.Context, db *sql.DB, dot *dotsql.DotSql) {
	_, err := dot.Exec(db, "delete-home-image", ctx.Params("id"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	ctx.Redirect("/admin?alert=Deleted+slider+image!#admin-homepage")
}

func AdminSliderOrder(ctx *macaron.Context, db *sql.DB, dot *dotsql.DotSql) {
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
	_, err = dot.Exec(db, "update-home-image-order", index, ctx.Params("id"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	ctx.Redirect("/admin#admin-homepage")
}
