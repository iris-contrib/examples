package main

import (
	"github.com/iris-contrib/plugin/cors"
	"gopkg.in/kataras/iris.v6"
)

func main() {
	app := iris.New()

	// acts like a wrapped router around the github.com/rs/cors net/http middleawre
	app.Plugins.Add(cors.Default())

	app.Get("/", func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "<b> Hello !</b>")
	})
	app.Listen(":80")
}
