package pages

import "gopkg.in/macaron.v1"

func Login(ctx *macaron.Context) {
	addStandardData(ctx.Data)
	ctx.HTMLSet(200, "base", "login")
}

func Logout(ctx *macaron.Context) {
	addStandardData(ctx.Data)
	//ctx.HTMLSet(200, "base", "login")
}
