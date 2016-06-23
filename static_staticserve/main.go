package main

import (
	"os"

	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/iris"
)

func main() {

	// Middleware
	iris.Use(recovery.New(os.Stderr))
	iris.Use(logger.New(iris.Logger))

	// Get theme
	iris.StaticServe("./resources", "/assets")

	// Router
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Append one of these to browser's address bar:\n/assets/js/jquery-2.1.1.js\n/assets/css/bootstrap.min.css")
	})

	iris.Listen(":8080")
}
