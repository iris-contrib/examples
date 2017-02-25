// Package main, simple example for context.Redirect, to use Redirect via {{ url }} use the context.RedirectTo, look at the template_engines examples.
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

	app.Get("/", func(ctx *iris.Context) {
		//ctx.Log(string(ctx.Request.URI.String())
		ctx.Redirect("/redirected")
	})

	app.Get("/redirected", func(ctx *iris.Context) {
		ctx.Writef("Hello, you have been redirected!")
	})

	app.Listen(":8080")
}
