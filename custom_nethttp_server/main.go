package main

import (
	"net/http"

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
		ctx.Writef("Hello from the server")
	})

	app.Get("/mypath", func(ctx *iris.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})
	// to use a custom server you have to call .Boot before server's listen
	app.Boot()

	// create our custom fasthttp server and assign the Handler/Router
	fsrv := &http.Server{Handler: app, Addr: ":8080"}
	fsrv.ListenAndServe()

	// now if you navigate to http://127.0.0.1:8080/mypath
}
