//Package main a basic and simple example on how to use handlebars with Iris

package main

import (
	"github.com/aymerick/raymond"

	"github.com/kataras/iris"
)

func main() {
	// set the template engine
	iris.Config.Render.Template.Engine = iris.HandlebarsEngine
	//iris.Config.Render.Template.Extensions = []string{".hbs"} // If you want to use the .hbs extension uncomment this line

	iris.Config.Render.Template.Layout = "layouts/layout.html" // default ""

	// NOTE:
	// the Iris' route framework {{url "my-routename" myparams}} and {{urlpath "my-routename" myparams}} are working like all other template engines,
	// so  avoid custom url and urlpath helpers.

	// optionaly set handlebars helpers by importing "github.com/aymerick/raymond" when you need to return and render html
	iris.Config.Render.Template.Handlebars.Helpers["boldme"] = func(input string) raymond.SafeString {
		return raymond.SafeString("<b> " + input + "</b>")
	}

	iris.Get("/", func(ctx *iris.Context) {
		// optionally, set a context  for the template
		if err := ctx.Render("home.html", map[string]interface{}{"Name": "Iris", "Type": "Web", "Path": "/"}); err != nil {
			println(err.Error())
		}
	})

	// remove the layout for a specific route
	iris.Get("/nolayout", func(ctx *iris.Context) {
		if err := ctx.Render("home.html", nil, iris.NoLayout); err != nil {
			println(err.Error())
		}
	})

	// set a layout for a party, .Layout should be BEFORE any Get or other Handle party's method
	my := iris.Party("/my").Layout("layouts/mylayout.html")
	{
		my.Get("/", func(ctx *iris.Context) {
			ctx.MustRender("home.html", map[string]interface{}{"Name": "Iris", "Type": "Web", "Path": "/my/"})
		})
		my.Get("/other", func(ctx *iris.Context) {
			ctx.MustRender("home.html", map[string]interface{}{"Name": "Iris", "Type": "Web", "Path": "/my/other"})
		})
	}

	iris.Listen(":8080")
}

/*
MORE DOCS CAN BE FOUND HERE: https://github.com/aymerick/raymond
*/
