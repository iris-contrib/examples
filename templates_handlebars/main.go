//Package main a basic and simple example on how to use handlebars with Iris

package main

import (
	"github.com/aymerick/raymond"

	"github.com/kataras/iris"
)

func main() {
	// set the template engine
	iris.Config.Render.Template.Engine = iris.HandlebarsEngine

	// optionaly set handlebars helpers by importing "github.com/aymerick/raymond" when you need to return and render html
	iris.Config.Render.Template.Handlebars.Helpers["boldme"] = func(input string) raymond.SafeString {
		return raymond.SafeString("<b> " + input + "</b>")
	}

	// NOTE:
	// the Iris' route framework {{url "my-routename" myparams}} and {{urlpath "my-routename" myparams}} are working like all other template engines,
	// so  avoid custom url and urlpath helpers.

	iris.Get("/", func(ctx *iris.Context) {
		// optionally, set a context  for the template
		mycontext := iris.Map{"Name": "Iris", "Type": "Web"}

		ctx.Render("home.html", mycontext)
	})
	iris.Listen(":8080")
}

/*
MORE DOCS CAN BE FOUND HERE: https://github.com/aymerick/raymond
*/
