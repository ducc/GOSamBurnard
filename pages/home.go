package pages

import "gopkg.in/macaron.v1"

func Home(ctx *macaron.Context) {
	addStandardData(ctx.Data)
    ctx.HTMLSet(200, "base", "index")
}