package main

import (
	"github.com/aymerick/raymond"
	"github.com/kataras/go-template/handlebars"
	"gopkg.in/kataras/iris.v6"
)

type mypage struct {
	Title   string
	Message string
}

func main() {
	// set the configuration for this template engine  (all template engines has its configuration)
	config := handlebars.DefaultConfig()
	config.Layout = "layouts/layout.html"
	config.Helpers["boldme"] = func(input string) raymond.SafeString {
		return raymond.SafeString("<b> " + input + "</b>")
	}

	// set the template engine
	iris.UseTemplate(handlebars.New(config)).Directory("./templates", ".html") // or .hbs , whatever you want

	iris.Get("/", func(ctx *iris.Context) {
		// optionally, set a context  for the template
		ctx.Render("home.html", map[string]interface{}{"Name": "Iris", "Type": "Web", "Path": "/"})

	})

	// remove the layout for a specific route using iris.NoLayout
	iris.Get("/nolayout", func(ctx *iris.Context) {
		if err := ctx.Render("home.html", nil, iris.RenderOptions{"layout": iris.NoLayout}); err != nil {
			ctx.Writef(err.Error())
		}
	})

	// set a layout for a party, .Layout should be BEFORE any Get or other Handle party's method
	my := iris.Party("/my").Layout("layouts/mylayout.html")
	{
		my.Get("/", func(ctx *iris.Context) {
			// .MustRender -> same as .Render but logs the error if any and return status 500 on client
			ctx.MustRender("home.html", map[string]interface{}{"Name": "Iris", "Type": "Web", "Path": "/my/"})
		})
		my.Get("/other", func(ctx *iris.Context) {
			ctx.MustRender("home.html", map[string]interface{}{"Name": "Iris", "Type": "Web", "Path": "/my/other"})
		})
	}

	iris.Listen(":8080")
}

// Note than you can see more handlebars examples syntax by navigating to https://github.com/aymerick/raymond
