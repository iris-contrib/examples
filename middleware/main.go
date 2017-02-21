package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	// register global middleware, you can pass more than one handler comma separated
	app.UseFunc(func(ctx *iris.Context) {
		println("(1)Global logger path: " + ctx.Path())
		ctx.Next()
	})

	// register a global structed iris.Handler as middleware
	myglobal := MyGlobalMiddlewareStructed{loggerId: "my logger id"}
	app.Use(myglobal)

	// register route's middleware
	app.Get("/home", func(ctx *iris.Context) {
		println("(1)HOME logger for /home")
		ctx.Next()
	}, func(ctx *iris.Context) {
		println("(2)HOME logger for /home")
		ctx.Next()
	}, func(ctx *iris.Context) {
		ctx.Writef("Hello from /home")
	})

	app.Listen(":8080")
}

// a silly example
type MyGlobalMiddlewareStructed struct {
	loggerId string
}

//Important staff, iris middleware must implement the iris.Handler interface which is:
func (m MyGlobalMiddlewareStructed) Serve(ctx *iris.Context) {
	println("Hello from the MyGlobalMiddlewareStructed")
	ctx.Next()
}
