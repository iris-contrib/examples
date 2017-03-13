package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	app.Config.Charset = "UTF-8" // this is the default, you don't have to set it manually

	myString := "this is just a simple string which you can already render with ctx.Write"

	app.Get("/", func(ctx *iris.Context) {
		ctx.Text(iris.StatusOK, myString)
	})

	app.Get("/alternative_1", func(ctx *iris.Context) {
		ctx.Render("text/plain", myString)
	})

	app.Get("/alternative_2", func(ctx *iris.Context) {
		ctx.RenderWithStatus(iris.StatusOK, "text/plain", myString)
	})

	app.Get("/alternative_3", func(ctx *iris.Context) {
		ctx.Render("text/plain", myString, iris.RenderOptions{"charset": "UTF-8"}) // default & global charset is UTF-8
	})

	app.Get("/alternative_4", func(ctx *iris.Context) {
		// logs if any error and sends http status '500 internal server error' to the client
		ctx.MustRender("text/plain", myString)
	})

	app.Listen(":8080")
}
