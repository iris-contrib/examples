package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Config.Render.Template.Layout = "layouts/layout.html" // default ""
	iris.Get("/", func(ctx *iris.Context) {
		s := iris.TemplateString("page1.html", nil)
		ctx.Write("The plain content of the template is: %s", s)
	})

	iris.Listen(":8080")
}
