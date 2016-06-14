package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

/*
With options:

errorLogger := logger.New(iris.Logger, logger.Options{
		EnableColors: false, //enable it to enable colors for all, disable colors by iris.Logger.ResetColors(), defaults to false
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
})

iris.Use(errorLogger)

With default options:

iris.Use(logger.New(iris.Logger))
*/
func main() {

	iris.Use(logger.New(iris.Logger))

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("hello")
	})

	iris.Get("/1", func(ctx *iris.Context) {
		ctx.Write("hello")
	})

	iris.Get("/2", func(ctx *iris.Context) {
		ctx.Write("hello")
	})

	// log http errors
	errorLogger := logger.New(iris.Logger)

	// yes we have options look at the logger.Options inside kataras/iris/middleware/logger.go
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.Write("My Custom 404 error page ")
	})
	//

	iris.Listen(":8080")

}

/* Book section: 'Logger' */
