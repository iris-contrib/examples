package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		app.Log(iris.DevMode, "/:not_found")
		ctx.HTML(iris.StatusNotFound, "<h1>/:not_found</h1><h2> Not Found Custom Message or Render a template | "+ctx.Path()+"</h2>")
	})

	static := app.Party("/static")
	{
		static.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
			app.Log(iris.DevMode, "/static:not_found")
			ctx.HTML(iris.StatusNotFound, "<h1>/static:not_found</h1><h2> Not Found Custom Message or Render a template</h2>")
		})
		// or static.StaticWeb("/", "./static")
		// or static.Get("/*file", iris.StripPrefix("/static", iris.StaticHandler("./static", false, true)))
		// or (depends on your needs)
		static.Get("/*file", static.StaticHandler("", "./static", false, true))
	}

	app.Listen(":8080")
}
