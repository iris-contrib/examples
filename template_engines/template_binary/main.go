package main

import (
	"github.com/kataras/go-template/html"
	"gopkg.in/kataras/iris.v6"
)

func main() {
	iris.Config.IsDevelopment = true // this will reload the templates on each request, defaults to false

	//$ go-bindata ./templates/...
	// templates are not used, you can delete the folder and run the example
	iris.UseTemplate(html.New()).Directory("./templates", ".html").Binary(Asset, AssetNames)

	iris.Get("/hi", hi)
	iris.Listen(":8080")
}

func hi(ctx *iris.Context) {
	ctx.MustRender("hi.html", struct{ Name string }{Name: "iris"})
}
