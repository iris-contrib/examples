package main

import (
	"github.com/kataras/iris"
)

type mypage struct {
	Title   string
	Message string
}

func main() {

	//optionally - before the load.
	//iris.Config.Render.Template.HTMLTemplate.Left = "${"  // Default is "{{"
	//iris.Config.Render.Template.HTMLTemplate.Right = "}" // Default is "}}"
	//iris.Config.Render.Template.HTMLTemplate.Funcs = template.FuncMap(...)

	//iris.Config.Render.Template.Directory = "templates" // Default is "templates"
	iris.Config.Render.Template.IsDevelopment = true // rebuild the templates on each refresh. Default is false
	//api.Config.Render.Template.Layout = "layout.html" // means: ./templates/layout.html.  Default is ""
	iris.Config.Render.Template.Gzip = true // Default is false

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("mypage.html", mypage{"My Page title", "Hello world!"}) //, "otherLayout" <- to override the layout
	})

	iris.Get("/hi_json", func(c *iris.Context) {
		c.JSON(200, iris.Map{
			"Name": "Iris",
			"Age":  2,
		})
	})

	iris.Listen(":8080")
}
