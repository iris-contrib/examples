package main

import "gopkg.in/kataras/iris.v6"

func main() {
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusNotFound, "<h1> Custom Not Found Message </h1>")
	})

	iris.Get("/howto", func(ctx *iris.Context) { ctx.Writef("Go to /static to view your static index.html") })

	iris.StaticWeb("/static", "./static")
	// or
	// iris.Get("/static/*file", iris.StaticHandler("/", "./static", false, false))
	iris.Listen(":8080")
}
