package main

import (
	"github.com/iris-contrib/template/markdown"
	"github.com/kataras/iris"
)

type mypage struct {
	Title   string
	Message string
}

func main() {

	iris.UseEngine(markdown.New()).Directory("./templates", ".md")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.MustRender("index.md", nil)
		// doesnt' supports any context binding, just pure markdown
	})

	iris.Listen(":8080")
}
