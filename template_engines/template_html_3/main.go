package main

import (
	"github.com/kataras/go-template/html"
	"gopkg.in/kataras/iris.v6"
)

func main() {
	// directory and extensions defaults to ./templates, .html for all template engines
	iris.UseTemplate(html.New(html.Config{Layout: "layouts/layout.html"}))

	iris.Get("/", func(ctx *iris.Context) {
		s := iris.TemplateString("page1.html", nil)
		ctx.Writef("The plain content of the template is: %s", s)
	})

	iris.Listen(":8080")
}
