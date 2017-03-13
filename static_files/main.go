package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

type page struct {
	Title string
}

func main() {

	app := iris.New()
	app.Adapt(
		iris.DevLogger(),
		httprouter.New(),
		view.HTML("./templates", ".html"),
	)

	app.OnError(iris.StatusForbidden, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusForbidden, "<h1> You are not allowed here </h1>")
	})
	// http://localhost:8080/css/bootstrap.min.css
	app.StaticWeb("/css", "./resources/css")
	// http://localhost:8080/js/jquery-2.1.1.js
	app.StaticWeb("/js", "./resources/js")

	app.Get("/", func(ctx *iris.Context) {
		ctx.MustRender("something.html", page{Title: "Home"})
	})

	app.Listen("localhost:8080")
}
