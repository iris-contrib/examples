package main

import "gopkg.in/kataras/iris.v5"

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
