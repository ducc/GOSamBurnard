package pages

import "gopkg.in/macaron.v1"

func Admin(ctx *macaron.Context) {
    ctx.HTMLSet(200, "base", "admin")
}