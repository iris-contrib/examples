package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recovery"
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
