package main

import (
	"github.com/kataras/iris"
)

func main() {
	api := iris.New()

	api.Config.Render.Template.Engine = iris.PongoEngine

	api.Get("/", func(ctx *iris.Context) {

		ctx.MustRender("index.html", map[string]interface{}{"username": "iris", "is_admin": true})

		/*
			//OR
				err := ctx.Render("index.html", pongo2.Context{"username": "iris", "is_admin": true})
				if err != nil {
					panic(err)
				}
		*/
	})

	api.Listen(":8080")
}
