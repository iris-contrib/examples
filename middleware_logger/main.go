package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/logger"
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
	})

	app.Use(customLogger)

	app.Get("/", func(ctx *iris.Context) {
		ctx.Writef("hello")
	})

	app.Get("/1", func(ctx *iris.Context) {
		ctx.Writef("hello")
	})

	app.Get("/2", func(ctx *iris.Context) {
		ctx.Writef("hello")
	})

	// log http errors
	errorLogger := logger.New()

	// yes we have options look at the logger.Options inside kataras/iris/middleware/logger.go
	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.Writef("My Custom 404 error page ")
	})
	//

	app.Listen(":8080")

}
