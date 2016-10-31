package main

import (
	"gopkg.in/kataras/iris.v4"
	"github.com/valyala/fasthttp"
)

func main() {
	api := iris.New()
	api.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello from the server")
	})

	api.Get("/mypath", func(ctx *iris.Context) {
		ctx.Write("Hello from the server on path /mypath")
	})
	// to use a custom server you have to call .Build after route, sessions, templates, websockets,ssh... before server's listen
	api.Build()

	// optionally, if you use plugins call:	api.Plugins.DoPreListen(api) before ListenAndServe
	// and api.Plugins.DoPostListen(api) after he ListenAndServe

	// create our custom fasthttp server and assign the Handler/Router
	fsrv := &fasthttp.Server{Handler: api.Router}
	fsrv.ListenAndServe(":8080")

	// now if you navigate to http://127.0.0.1:8080/mypath
}
