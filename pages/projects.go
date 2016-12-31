package pages

import "gopkg.in/macaron.v1"

func Projects(ctx *macaron.Context) {
	addStandardData(ctx.Data, "projects")
	ctx.HTMLSet(200, "base", "projects")
}

func Project(ctx *macaron.Context) {
	addStandardData(ctx.Data, "projects")
	ctx.HTMLSet(200, "base", "project")
}
