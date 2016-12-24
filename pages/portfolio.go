package pages

import (
    "gopkg.in/macaron.v1"
)

type portfolioItem struct {
	id          uint
	image       string
	title       string
	description string
}

func Portfolio(ctx *macaron.Context) {
	addStandardData(ctx.Data)
    images := make([]portfolioItem, 0)
    images = append(images, portfolioItem{
        id: 1,
        image: "https://upload.wikimedia.org/wikipedia/commons/thumb/d/d4/CH_cow_2_cropped.jpg/250px-CH_cow_2_cropped.jpg",
        title: "Cow 1 !!",
        description: "this is my favorite cow",
    }, portfolioItem{
        id: 2,
        image: "https://i.ytimg.com/vi/7FriQgP-3DM/maxresdefault.jpg",
        title: "GET SHREK'D",
        description: "this is my favorite cow",
    }, portfolioItem{
        id: 3,
        image: "https://upload.wikimedia.org/wikipedia/commons/thumb/d/d4/CH_cow_2_cropped.jpg/250px-CH_cow_2_cropped.jpg",
        title: "Cow 3 !!",
        description: "this is my favorite cow",
    })
    ctx.Data["images"] = images
    ctx.HTMLSet(200, "base", "portfolio")
}