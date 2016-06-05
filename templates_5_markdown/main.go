package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
)

func main() {
	// Markdown engine doesn't supports Layout and context binding
	iris.Config().Render.Template.Engine = config.MarkdownEngine
	iris.Config().Render.Template.Extensions = []string{".md"}

	iris.Get("/", func(ctx *iris.Context) {

		err := ctx.Render("index.md", nil) // doesnt' supports any context binding, just pure markdown
		if err != nil {
			panic(err)
		}
	})

	iris.Listen(":8080")
}
