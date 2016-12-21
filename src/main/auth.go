package main

import "gopkg.in/macaron.v1"

func login(ctx *macaron.Context) {
	ctx.HTMLSet(200, "base", "login")
}

func logout(ctx *macaron.Context) {
	//ctx.HTMLSet(200, "base", "login")
}