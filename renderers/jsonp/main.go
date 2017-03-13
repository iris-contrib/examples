package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

type myjson struct {
	Name string `json:"name"`
}

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	app.Get("/", func(ctx *iris.Context) {
		ctx.JSONP(iris.StatusOK, "callbackName", iris.Map{"name": "iris"})
	})

	app.Get("/alternative_1", func(ctx *iris.Context) {
		ctx.JSONP(iris.StatusOK, "callbackName", myjson{Name: "iris"})
	})

	app.Get("/alternative_2", func(ctx *iris.Context) {
		ctx.Render("application/javascript", myjson{Name: "iris"}, iris.RenderOptions{"callback": "callbackName"})
	})

	app.Get("/alternative_3", func(ctx *iris.Context) {
		ctx.RenderWithStatus(iris.StatusOK, "application/javascript", myjson{Name: "iris"}, iris.RenderOptions{"callback": "callbackName"})
	})

	app.Get("/alternative_4", func(ctx *iris.Context) {
		// logs if any error and sends http status '500 internal server error' to the client
		ctx.MustRender("application/javascript", myjson{Name: "iris"}, iris.RenderOptions{"callback": "callbackName", "charset": "UTF-8"}) // UTF-8 is the default.
	})

	app.Listen(":8080")
}
