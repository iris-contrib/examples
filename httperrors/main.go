package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view" // optional
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())
	// load the view templates from ./templates folder and .html extension
	// use the standard html/template syntax
	app.Adapt(view.HTML("./templates", ".html"))

	app.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Writef(iris.StatusText(iris.StatusInternalServerError)) // Outputs: Internal Server Error
		ctx.SetStatusCode(iris.StatusInternalServerError)           // 500

		println("http status: 500 happened!")
	})

	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.RenderWithStatus(iris.StatusNotFound, "errors/404.html", nil)
	})

	// emit the errors to test them
	app.Get("/500", func(ctx *iris.Context) {
		ctx.EmitError(iris.StatusInternalServerError) // ctx.Panic()
	})

	app.Get("/404", func(ctx *iris.Context) {
		ctx.EmitError(iris.StatusNotFound) // ctx.NotFound()
	})

	// navigate to localhost:8080/dsajdsada and you will see the custom http error 404
	// or navigate to localhost:8080/404 and localhost:8080/500 to emit the errors manually

	users := app.Party("/users")
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
			ctx.Writef("Hello from Profile with ID: %s", ctx.Param("id"))
		})
	}

	// navigate to localhost:8080/users/dsajdsada and you will see the custom http error 404
	// or navigate to localhost:8080/users/404 and localhost:8080/users/500 to emit the errors manually

	app.Listen(":8080")

}
