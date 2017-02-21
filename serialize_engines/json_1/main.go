package main

import "gopkg.in/kataras/iris.v6"

type myjson struct {
	Name string `json:"name"`
}

func main() {

	iris.Get("/", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, iris.Map{"name": "iris"})
	})

	iris.Get("/alternative_1", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, myjson{Name: "iris"})
	})

	iris.Get("/alternative_2", func(ctx *iris.Context) {
		ctx.Render("application/json", myjson{Name: "iris"})
	})

	iris.Get("/alternative_3", func(ctx *iris.Context) {
		ctx.RenderWithStatus(iris.StatusOK, "application/json", myjson{Name: "iris"})
	})

	iris.Get("/alternative_4", func(ctx *iris.Context) {
		ctx.Render("application/json", myjson{Name: "iris"}, iris.RenderOptions{"charset": "UTF-8"}) // UTF-8 is the default.
	})

	iris.Get("/alternative_5", func(ctx *iris.Context) {
		// logs if any error and sends http status '500 internal server error' to the client
		ctx.MustRender("application/json", myjson{Name: "iris"}, iris.RenderOptions{"charset": "UTF-8"}) // UTF-8 is the default.
	})

	iris.Listen(":8080")
}
