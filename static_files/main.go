package main

import (
	"github.com/kataras/go-template/html"
	"gopkg.in/kataras/iris.v6"
)

type page struct {
	Title string
}

func main() {
	iris.UseTemplate(html.New()).Directory("./templates/web/default", ".html")
	iris.OnError(iris.StatusForbidden, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusForbidden, "<h1> You are not allowed here </h1>")
	})
	// http://localhost:8080/css/bootstrap.min.css
	iris.StaticWeb("/css", "./resources/css")
	// http://localhost:8080/js/jquery-2.1.1.js
	iris.StaticWeb("/js", "./resources/js")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.MustRender("something.html", page{Title: "Home"})
	})

	iris.Listen(":8080")
}
