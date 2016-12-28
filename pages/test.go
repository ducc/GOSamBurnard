package pages

import "gopkg.in/macaron.v1"

func Test(ctx *macaron.Context) {
	addStandardData(ctx.Data)
    ctx.HTMLSet(200, "base", "test")
}

func TestSubmit() {

}