package main

import "github.com/kataras/iris"

func main() {
	err := iris.Favicon("./static/favicons/iris_favicon_32_32.ico")
	// This will serve the ./static/favicons/iris_favicon_32_32.ico to: localhost:8080/favicon.ico

	// err := iris.Favicon("./static/favicons/iris_favicon_32_32.ico", "/favicon_32_32.ico")
	// This will serve the ./static/favicons/iris_favicon_32_32.ico to: localhost:8080/favicon_32_32.ico

	if err != nil {
		iris.Logger().Panicf("Error when trying to set static favicon %s", err.Error())
	}

	iris.Get("/", func(ctx *iris.Context) {
		ctx.WriteHTML(iris.StatusOK, "You should see the favicon now at the side of your browser, if not please refresh or clear the browser's cache.")
	})

	iris.Listen(":8080")
}
