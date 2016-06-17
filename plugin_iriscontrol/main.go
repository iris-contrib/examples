package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/plugin/iriscontrol"
)

func main() {

	iris.Plugins.Add(iriscontrol.New(9090, map[string]string{
		"1":             "1",
		"irisusername2": "irispassowrd2",
	}))

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Root path from  server")
	})

	iris.Get("/something", func(ctx *iris.Context) {
		ctx.Write("Something path from server")
	})

	iris.Listen(":8080")
}
