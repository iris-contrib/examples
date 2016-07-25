package main

import "github.com/kataras/iris"

func firstMiddleware(ctx *iris.Context) {
	ctx.Write("1. This is the first middleware, before any of route's handlers \n")
	ctx.Next()
}

func secondMiddleware(ctx *iris.Context) {
	ctx.Write("2. This is the second middleware, before the / main handler \n")
	ctx.Next()
}

func thirdMiddleware(ctx *iris.Context) {
	ctx.Write("3. This is the 3rd middleware, after the main handler \n")
	ctx.Next()
}

func lastAlwaysMiddleware(ctx *iris.Context) {
	ctx.Write("4. This is the ALWAYS LAST Handler \n")
}

func main() {

	iris.UseFunc(firstMiddleware)
	iris.DoneFunc(lastAlwaysMiddleware)

	iris.Get("/", secondMiddleware, func(ctx *iris.Context) {
		ctx.Write("Hello from / \n")
		ctx.Next() // .Next because we 're using the third middleware after that, and lastAlwaysMiddleware also
	}, thirdMiddleware)

	// for parties look the 'middleware_3' example

	iris.Listen(":8080")

}
