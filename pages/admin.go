package pages

import (
	"database/sql"
	"github.com/gchaincl/dotsql"
	"gopkg.in/macaron.v1"
	"log"
)

func Admin(ctx *macaron.Context, db *sql.DB, dot *dotsql.DotSql) {
	addStandardData(ctx.Data)
	var err error
	ctx.Data["slider_items"], err = loadAndSortSliderItems(db, dot)
	ctx.Data["portfolio_images"], err = loadAndSortPortfolioItems(db, dot)
	ctx.Data["about_text"], ctx.Data["contact_text"] = about, contact
	ctx.Data["admin_social_accounts"], err = getSocialAccounts(db, dot)
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx.HTMLSet(200, "base", "admin")
}