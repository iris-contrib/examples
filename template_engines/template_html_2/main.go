package main

import (
	"github.com/kataras/go-template/html"
	"gopkg.in/kataras/iris.v6"
)

func main() {
	// directory and extensions defaults to ./templates, .html for all template engines
	iris.UseTemplate(html.New(html.Config{Layout: "layouts/layout.html"}))
	//iris.Config.Render.Template.Gzip = true
	iris.Get("/", func(ctx *iris.Context) {
		if err := ctx.Render("page1.html", nil); err != nil {
			println(err.Error())
		}
	})

	// remove the layout for a specific route
	iris.Get("/nolayout", func(ctx *iris.Context) {
		if err := ctx.Render("page1.html", nil, iris.RenderOptions{"layout": iris.NoLayout}); err != nil {
			println(err.Error())
		}
	})

	// set a layout for a party, .Layout should be BEFORE any Get or other Handle party's method
	my := iris.Party("/my").Layout("layouts/mylayout.html")
	{
		my.Get("/", func(ctx *iris.Context) {
			ctx.MustRender("page1.html", nil)
		})
		my.Get("/other", func(ctx *iris.Context) {
			ctx.MustRender("page1.html", nil)
		})
	}

	iris.Listen(":8080")
}
