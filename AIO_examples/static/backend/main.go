package main

import (
	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
)

func main() {
	// set the configuration
	// we need this only to set the custom 404 page, otherwise we can skip that:
	iris.Config.Render.Template.Directory = "../frontend/templates"
	// set the custom error(s)
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.MustRender("404.html", nil)
	})

	// set the middleware(s)
	iris.Use(logger.New(iris.Logger))

	// if you want to publish just a static website you don't have to set any routes
	// Iris has one-line method to do that:
	iris.StaticWeb("/", "../frontend/webstatic", 0)

	// start the server
	iris.Listen("127.0.0.1:80")
}
