package main

import (
	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"gopkg.in/kataras/iris.v5"
)

func main() {

	// Middleware
	iris.Use(recovery.New())
	iris.Use(logger.New())

	// Get theme
	iris.StaticServe("./resources", "/assets")

	// Router
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Append one of these to browser's address bar:\n/assets/js/jquery-2.1.1.js\n/assets/css/bootstrap.min.css")
	})

	iris.Listen(":8080")
}
