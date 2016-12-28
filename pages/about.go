package pages

import "gopkg.in/macaron.v1"

func About(ctx *macaron.Context) {
	addStandardData(ctx.Data, "about")
    ctx.HTMLSet(200, "base", "about")
}