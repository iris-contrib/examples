package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Config().Render.Template.Layout = "layouts/layout.html" // default ""
	//iris.Config().Render.Template.Gzip = true
	iris.Get("/", func(ctx *iris.Context) {
		if err := ctx.Render("page1.html", nil); err != nil {
			println(err.Error())
		}
	})

	iris.Listen(":8080")
}
