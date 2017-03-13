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

	myData := []byte("some binary data or a program here which will not be a simple string at the production")

	app.Get("/", func(ctx *iris.Context) {
		ctx.Data(iris.StatusOK, myData)
	})

	app.Get("/alternative_1", func(ctx *iris.Context) {
		ctx.Render("application/octet-stream", myData)
	})

	app.Get("/alternative_2", func(ctx *iris.Context) {
		ctx.RenderWithStatus(iris.StatusOK, "application/octet-stream", myData)
	})

	app.Get("/alternative_3", func(ctx *iris.Context) {
		ctx.Render("application/octet-stream", myData, iris.RenderOptions{"gzip": true}) // gzip is false by default
	})

	app.Get("/alternative_4", func(ctx *iris.Context) {
		// logs if any error and sends http status '500 internal server error' to the client
		ctx.MustRender("application/octet-stream", myData)
	})

	app.Listen(":8080")
}
