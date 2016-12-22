package pages

import "gopkg.in/macaron.v1"

func Home(ctx *macaron.Context) {
    ctx.HTMLSet(200, "base", "index")
}