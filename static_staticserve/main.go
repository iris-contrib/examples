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
	iris.Use(logger.New(iris.Logger()))

	// Get theme
	iris.StaticServe("./resources", "/assets")

	// Router
	iris.Get("/", func(response *iris.Context) {
		response.Write("Test")
	})

	iris.Listen(":8080")
}
