package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"gopkg.in/kataras/iris.v6/middleware/logger"
)

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())
	// we need this only to set the custom 404 page, otherwise we can skip that:
	// set the template engine
	app.Adapt(view.HTML("../frontend/templates", ".html"))

	// set the custom error(s)
	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.MustRender("404.html", nil)
	})

	// set the middleware(s)
	app.Use(logger.New())

	// if you want to publish just a static website you don't have to set any routes
	// Iris has one-line method to do that:
	app.StaticWeb("/", "../frontend/webstatic")

	// start the server
	app.Listen("127.0.0.1:8080")
}
