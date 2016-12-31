package main

import "github.com/kataras/iris"

// if your ide cannot find the ./static folder try to build that program and after execute it
// or try to download & run this example via LiteIDE.
func main() {
	iris.Favicon("./static/favicons/iris_favicon_32_32.ico")
	// This will serve the ./static/favicons/iris_favicon_32_32.ico to: localhost:8080/favicon.ico

	// iris.Favicon("./static/favicons/iris_favicon_32_32.ico", "/favicon_32_32.ico")
	// This will serve the ./static/favicons/iris_favicon_32_32.ico to: localhost:8080/favicon_32_32.ico

	iris.Get("/", func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "You should see the favicon now at the side of your browser, if not please refresh or clear the browser's cache.")
	})

	iris.Listen(":8080")
}
