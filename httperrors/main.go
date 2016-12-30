package main

import (
	"gopkg.in/kataras/iris.v5"
)

func main() {

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Write(iris.StatusText(iris.StatusInternalServerError)) // Outputs: Internal Server Error
		ctx.SetStatusCode(iris.StatusInternalServerError)          // 500

		ctx.Log("http status: 500 happened!\n")
	})

	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Write(iris.StatusText(iris.StatusNotFound)) // Outputs: Not Found
		ctx.SetStatusCode(iris.StatusNotFound)          // 404

		ctx.Log("http status: 404 happened!\n")
	})

	// emit the errors to test them
	iris.Get("/500", func(ctx *iris.Context) {
		ctx.EmitError(iris.StatusInternalServerError) // ctx.Panic()
	})

	iris.Get("/404", func(ctx *iris.Context) {
		ctx.EmitError(iris.StatusNotFound) // ctx.NotFound()
	})

	// navigate to localhost:8080/dsajdsada and you will see the custom http error 404
	// or navigate to localhost:8080/404 and localhost:8080/500 to emit the errors manually

	users := iris.Party("/users")
	{
		users.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
			ctx.WriteString("This is a Not Found error for /users path prefix")
		})
		users.Get("/500", func(ctx *iris.Context) {
			ctx.EmitError(iris.StatusInternalServerError) // ctx.Panic()
		})

		users.Get("/404", func(ctx *iris.Context) {
			ctx.EmitError(iris.StatusNotFound) // ctx.NotFound()
		})

		users.Get("/profile/:id", func(ctx *iris.Context) {
			ctx.Write("Hello from Profile with ID: %s", ctx.Param("id"))
		})
	}

	// navigate to localhost:8080/users/dsajdsada and you will see the custom http error 404
	// or navigate to localhost:8080/users/404 and localhost:8080/users/500 to emit the errors manually

	iris.Listen(":8080")

}
