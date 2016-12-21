package main

import "gopkg.in/macaron.v1"

func admin(ctx *macaron.Context) {
    ctx.HTMLSet(200, "base", "admin")
}