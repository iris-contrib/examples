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

	// Let's party
	// A 'Party/Group of routes and APIs'
	// shares the same path prefix
	// the same middleware (second parameter)
	// and optionally the same template layout (admin.Layout("/layouts/admin.layout.html"))
	//
	// the same syntax for subdomains and wildcard subdomains too
	admin := app.Party("/admin")
	{
		// add a silly middleware
		admin.UseFunc(func(ctx *iris.Context) {
			//your authentication logic here...
			println("from ", ctx.Path())
			authorized := true
			if authorized {
				ctx.Next()
			} else {
				ctx.Text(iris.StatusUnauthorized, ctx.Path()+" is not authorized for you")
			}

		})
		admin.Get("/", func(ctx *iris.Context) {
			ctx.Writef("Hello World")
			ctx.Writef("from /admin/ or /admin if you pathcorrection on")
		})
		admin.Get("/dashboard", func(ctx *iris.Context) {
			ctx.Writef("/admin/dashboard")
		})
		admin.Delete("/delete/:userId", func(ctx *iris.Context) {
			ctx.Writef("admin/delete/%s", ctx.Param("userId"))
		})
	}

	beta := admin.Party("/beta")
	beta.Get("/hey", func(ctx *iris.Context) { ctx.Writef("hey from /admin/beta/hey") })

	//for subdomains goto: ../subdomains_1/main.go

	app.Listen(":8080")

}
