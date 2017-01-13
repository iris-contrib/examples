package main

import (
	"github.com/iris-contrib/plugin/gorillamux"
	"github.com/kataras/iris"
)

func main() {
	iris.Plugins.Add(gorillamux.New())

	// CUSTOM HTTP ERRORS ARE SUPPORTED
	// NOTE: Gorilla mux allows customization only on StatusNotFound(404)
	// Iris allows for everything, so you can register any other custom http error
	// but you have to call it manually from ctx.EmitError(status_code) // 500 for example
	// this will work because it's StatusNotFound:
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusNotFound, "<h1> CUSTOM NOT FOUND ERROR PAGE </h1>")
	})

	// GLOBAL/PARTY MIDDLEWARE ARE SUPPORTED
	iris.UseFunc(func(ctx *iris.Context) {
		println("Request: " + ctx.Path())
		ctx.Next()
	})

	// http://mydomain.com
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello from index")
	})

	/// -------------------------------------- IMPORTANT --------------------------------------
	/// GORILLA MUX PARAMETERS(regexp) ARE SUPPORTED
	/// http://mydomain.com/api/users/42
	/// ---------------------------------------------------------------------------------------
	iris.Get("/api/users/{userid:[0-9]+}", func(ctx *iris.Context) {
		ctx.Writef("User with id: %s", ctx.Param("userid"))
	})

	// PER-ROUTE MIDDLEWARE ARE SUPPORTED
	// http://mydomain.com/other
	iris.Get("/other", func(ctx *iris.Context) {
		ctx.Writef("/other 1 middleware \n")
		ctx.Next()
	}, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "<b>Hello from /other</b>")
	})

	// SUBDOMAINS ARE SUPPORTED
	// http://admin.mydomain.com
	iris.Party("admin.").Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello from admin. subdomain!")
	})

	// WILDCARD SUBDOMAINS ARE SUPPORTED
	// http://api.mydomain.com/hi
	// http://admin.mydomain.com/hi
	// http://x.mydomain.com/hi
	// [depends on your host configuration,
	// you will see an example(win) outside of this folder].
	iris.Party("*.").Get("/hi", func(ctx *iris.Context) {
		ctx.Writef("Hello from wildcard subdomain: %s", ctx.Subdomain())
	})

	// DOMAIN NAMING IS SUPPORTED
	// Custom domain is totally optionally, you can still use `iris.Listen(":8080") or any host` of course.
	iris.Listen("mydomain.com")
	// iris.Listen(":80")
}

/* HOSTS FILE LINES TO RUN THIS EXAMPLE:

127.0.0.1		mydomain.com
127.0.0.1		admin.mydomain.com
127.0.0.1		api.mydomain.com
127.0.0.1		x.mydomain.com

*/
