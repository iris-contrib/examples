package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func firstMiddleware(ctx *iris.Context) {
	ctx.Writef("1. This is the first middleware, before any of route's handlers \n")
	ctx.Next()
}

func secondMiddleware(ctx *iris.Context) {
	ctx.Writef("2. This is the second middleware, before the / main handler \n")
	ctx.Next()
}

func thirdMiddleware(ctx *iris.Context) {
	ctx.Writef("3. This is the 3rd middleware, after the main handler \n")
	ctx.Next()
}

func lastAlwaysMiddleware(ctx *iris.Context) {
	ctx.Writef("4. This is the ALWAYS LAST Handler \n")
}

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	app.UseFunc(firstMiddleware)
	app.DoneFunc(lastAlwaysMiddleware)

	app.Get("/", secondMiddleware, func(ctx *iris.Context) {
		ctx.Writef("Hello from %s \n", ctx.Path())
		ctx.Next() // .Next because we 're using the third middleware after that, and lastAlwaysMiddleware also
	}, thirdMiddleware)

	// for parties look the 'middleware_3' example

	app.Listen(":8080")

}
