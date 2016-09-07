package main

import (
	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
)

func main() {

	// we need this only to set the custom 404 page, otherwise we can skip that:
	// set the template engine
	iris.UseTemplate(html.New()).Directory("../frontend/templates", ".html")
	// set the custom error(s)
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.MustRender("404.html", nil)
	})

	// set the middleware(s)
	iris.Use(logger.New())

	// if you want to publish just a static website you don't have to set any routes
	// Iris has one-line method to do that:
	iris.StaticWeb("/", "../frontend/webstatic", 0)

	// start the server
	iris.Listen("127.0.0.1:8080")
}
