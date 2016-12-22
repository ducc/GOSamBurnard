package main

import (
    "gopkg.in/macaron.v1"
)

func portfolio(ctx *macaron.Context) {
    images := make([]portfolioItem, 0)
    images = append(images, portfolioItem{
        id: 1,
        image: "https://upload.wikimedia.org/wikipedia/commons/thumb/d/d4/CH_cow_2_cropped.jpg/250px-CH_cow_2_cropped.jpg",
        title: "Cow !!",
        description: "this is my favorite cow",
    }, portfolioItem{
        id: 1,
        image: "https://upload.wikimedia.org/wikipedia/commons/thumb/d/d4/CH_cow_2_cropped.jpg/250px-CH_cow_2_cropped.jpg",
        title: "Cow !!",
        description: "this is my favorite cow",
    }, portfolioItem{
        id: 1,
        image: "https://upload.wikimedia.org/wikipedia/commons/thumb/d/d4/CH_cow_2_cropped.jpg/250px-CH_cow_2_cropped.jpg",
        title: "Cow !!",
        description: "this is my favorite cow",
    })
    ctx.Data["images"] = images
    ctx.HTMLSet(200, "base", "portfolio")
}