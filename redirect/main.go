// Package main, simple example for context.Redirect, to use Redirect via {{ url }} use the context.RedirectTo, look at the template_engines examples.
package main

import "gopkg.in/kataras/iris.v4"

func main() {

	iris.Get("/", func(ctx *iris.Context) {
		//ctx.Log(string(ctx.RequestCtx.URI().FullURI()))
		ctx.Redirect("/redirected")
	})

	iris.Get("/redirected", func(ctx *iris.Context) {
		ctx.Write("Hello, you have been redirected!")
	})

	iris.Listen(":8080")
}
