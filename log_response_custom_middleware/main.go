package main

import (
	"fmt"
	"time"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	app.UseFunc(responseLogger) // global middleware, catch all routes

	app.Get("/", func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "<h1> Hello from / </h1>")
	})

	app.Get("/home", func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "<h1> Hello from /home</h1>")
	})

	app.Listen(":8080")
}

func responseLogger(ctx *iris.Context) {
	// optionally: use recorder in order to take the body written by the main handler
	w := ctx.Recorder()
	// process the actual request first, we don't want to have delays
	// so execute the next(which is the main) handler first.
	ctx.Next()

	date := time.Now().Format("01/02 - 15:04:05")
	fmt.Printf("%s\n%s", date, w.Body())
}
