package main

import (
	"github.com/kataras/iris"
)

func main() {

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Write(iris.StatusText(iris.StatusInternalServerError)) // Outputs: Internal Server Error
		ctx.SetStatusCode(iris.StatusInternalServerError)          // 500

		iris.Logger.Dangerf("http status: 500 happened!\n")
	})

	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Write(iris.StatusText(iris.StatusNotFound)) // Outputs: Not Found
		ctx.SetStatusCode(iris.StatusNotFound)          // 404

		iris.Logger.Infof("http status: 404 happened!\n")
	})

	// emit the errors to test them
	iris.Get("/500", func(ctx *iris.Context) {
		ctx.EmitError(iris.StatusInternalServerError) // ctx.Panic()
	})

	iris.Get("/404", func(ctx *iris.Context) {
		ctx.EmitError(iris.StatusNotFound) // ctx.NotFound()
	})

	iris.Listen(":8080")

}
