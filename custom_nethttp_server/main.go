package main

import (
	"gopkg.in/kataras/iris.v6"
	"net/http"
)

func main() {
	api := iris.New()
	api.Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello from the server")
	})

	api.Get("/mypath", func(ctx *iris.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})
	// to use a custom server you have to call .Build after route, sessions, templates, websockets,ssh... before server's listen
	api.Build()

	// optionally, if you use plugins call:	api.Plugins.DoPreListen(api) before ListenAndServe
	// and api.Plugins.DoPostListen(api) after the ListenAndServe

	// create our custom fasthttp server and assign the Handler/Router
	fsrv := &http.Server{Handler: api.Router, Addr: ":8080"}
	fsrv.ListenAndServe()

	// now if you navigate to http://127.0.0.1:8080/mypath
}
