package main

import (
	"fmt"

	"github.com/kataras/iris"
)

func main() {
	// first way:
	// simple way for simple things
	// PreListen before a station is listening ( iris.Listen/ListenTLS/ListenLETSENCRYPT/ListenUNIX/Serve...)
	iris.Plugins.PreListen(func(s *iris.Framework) {
		for _, route := range s.Lookups() {
			fmt.Printf("Func: Route Method: %s | Subdomain %s | Path: %s is going to be registed with %d handler(s). \n", route.Method(), route.Subdomain(), route.Path(), len(route.Middleware()))
		}
	})

	// second way:
	// structured way for more things
	plugin := myPlugin{}
	iris.Plugins.Add(plugin)

	iris.Get("/first_route", aHandler)

	iris.Post("/second_route", aHandler)

	iris.Put("/third_route", aHandler)

	iris.Get("/fourth_route", aHandler)

	iris.Listen(":8080")
}

func aHandler(ctx *iris.Context) {
	ctx.Write("Hello from: %s", ctx.PathString())
}

type myPlugin struct{}

// PostListen after a station is listening ( iris.Listen/ListenTLS/ListenLETSENCRYPT/ListenUNIX/Serve...)
func (pl myPlugin) PostListen(s *iris.Framework) {
	fmt.Printf("myPlugin: server is listening on host: %s", s.Config.VHost)
}

//list:
/*
	Activate(iris.PluginContainer)
	GetName() string
	GetDescription() string
	PreListen(*iris.Framework)
	PostListen(*iris.Framework)
	PreClose(*iris.Framework)
	PreDownload(thePlugin iris.Plugin, downloadUrl string)
	// for custom events use go-events, https://github.com/kataras/go-events
*/
