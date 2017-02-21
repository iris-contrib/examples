package main

import (
	"github.com/kataras/go-template/markdown"
	"gopkg.in/kataras/iris.v6"
)

type mypage struct {
	Title   string
	Message string
}

func main() {
	iris.Config.Gzip = true // this will use gzip compression to serve the templates, defaults to false [ from 5k to 2.2k ]
	iris.UseTemplate(markdown.New()).Directory("./templates", ".md")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.MustRender("index.md", nil)
		// doesnt' supports any context binding, just pure markdown
	})

	iris.Listen(":8080")
}
