package main

import (
	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
)

type page struct {
	Title string
}

func main() {
	iris.UseTemplate(html.New()).Directory("./templates/web/default", ".html")
	iris.OnError(iris.StatusForbidden, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusForbidden, "<h1> You are not allowed here </h1>")
	})
	iris.StaticWeb("/css", "./resources/css")
	iris.StaticWeb("/js", "./resources/js")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.MustRender("something.html", page{Title: "Home"})
	})

	iris.Listen(":8080")
}
