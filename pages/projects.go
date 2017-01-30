package pages

import (
	"gopkg.in/macaron.v1"
	"github.com/gchaincl/dotsql"
	"database/sql"
)

func Projects(ctx *macaron.Context, db *sql.DB, dot *dotsql.DotSql) {
	addStandardData(ctx.Data, db, dot, "projects")
	ctx.HTMLSet(200, "base", "soon")
}

/*func Project(ctx *macaron.Context) {
	addStandardData(ctx.Data, "projects")
	ctx.HTMLSet(200, "base", "project")
}*/
