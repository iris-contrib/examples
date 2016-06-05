package main

import (
	"fmt"

	"github.com/kataras/iris"
)

func main() {
	// first way:
	// simple way for simple things
	// PreHandle before a route is registed ( iris.Get/Post...)
	iris.Plugins().PreHandle(func(route iris.IRoute) {
		fmt.Printf("Func: Route Method: %s and Path: %s is going to be registed with %d handler(s). \n", route.GetMethod(), route.GetPath(), len(route.GetMiddleware()))

	})

	// second way:
	// structured way for more things
	plugin := myPlugin{}
	iris.Plugins().Add(plugin)

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

// PostHandle after a route is registed ( iris.Get/Post...)
func (pl myPlugin) PostHandle(route iris.IRoute) {
	fmt.Printf("myPlugin: Route Method: %s and Path: %s registed with %d handler(s). \n", route.GetMethod(), route.GetPath(), len(route.GetMiddleware()))
}

// after iris.Listen
func (pl myPlugin) PostListen(station *iris.Iris) {
	fmt.Printf("myPlugin: Server is succesfuly running on address: %s. \n ", station.Server().Config.ListeningAddr)
}
