package main

import "github.com/kataras/iris"

func main() {

	iris.Config().Render.Template.Engine = iris.AmberEngine
	iris.Config().Render.Template.Extensions = []string{".amber"} // this is optionally, you can just leave it to default which is .html

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("basic.amber", map[string]string{"Name": "iris"})

	})
	println("Server is running at: 8080")
	iris.Listen(":8080")
}
