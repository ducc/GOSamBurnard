package pages

import "gopkg.in/macaron.v1"

func Login(ctx *macaron.Context) {
	ctx.HTMLSet(200, "base", "login")
}

func Logout(ctx *macaron.Context) {
	//ctx.HTMLSet(200, "base", "login")
}