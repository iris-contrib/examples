package main

import (
	"context"
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

	app.Adapt(iris.EventPolicy{
		// Interrupt Event means when control+C pressed on terminal.
		Interrupted: func(*iris.Framework) {
			println("control+C pressed, do your external cleanup here!")

			// when os.Interrupt signal is fired the body of this function will be
			// fired,
			// you're responsible for closing the server with app.Shutdown(...)

			// if that event is not registered then the framework
			// will gracefully close the server for you.
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			app.Shutdown(ctx)
		},
	})

	app.Get("/", func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "<h1>Hello from index!</h1>")
	})

	ln, err := iris.TCP4(":8080")
	if err != nil {
		panic(err)
	}
	app.Serve(ln)

}
