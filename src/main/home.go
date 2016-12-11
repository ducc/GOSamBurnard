package main

import "gopkg.in/macaron.v1"

func home(ctx *macaron.Context) {
    ctx.HTMLSet(200, "base", "index")
}