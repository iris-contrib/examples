package main

import (
	"github.com/kataras/iris"
)

type page struct {
	Title string
}

func main() {
	iris.Config.Render.Template.Directory = "templates\\web\\default"

	iris.OnError(iris.StatusForbidden, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusForbidden, "<h1> You are not allowed here </h1>")
	})
	iris.Static("/css", "./resources/css", 1)
	iris.Static("/js", "./resources/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.MustRender("something.html", page{Title: "Home"})
	})

	iris.Listen(":8080")
}
