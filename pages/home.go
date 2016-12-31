package pages

import "gopkg.in/macaron.v1"

func Home(ctx *macaron.Context) {
	addStandardData(ctx.Data, "home")
	ctx.HTMLSet(200, "base", "index")
}
