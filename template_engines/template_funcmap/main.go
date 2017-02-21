package main

import (
	// we need that ONLY to render the trusted template.HTML from a template func, golang doesn't supports type alias yet.
	"html/template"
	"time"

	"gopkg.in/kataras/iris.v6"
)

func main() {
	iris.Config.IsDevelopment = true
	iris.Config.EnablePathEscape = true

	// UseTemplateFunc sets or replaces a TemplateFunc from the shared available TemplateFuncMap
	// defaults are the iris.URL and iris.Path, all the template engines supports the following:
	// {{ url "mynamedroute" "pathParameter_ifneeded"} }
	// {{ urlpath "mynamedroute" "pathParameter_ifneeded" }}
	// {{ render "header.html" }}
	// {{ render_r "header.html" }} // partial relative path to current page
	// {{ yield }}
	// {{ current }}
	// see other template_engines example folders for more.
	iris.UseTemplateFunc("bold", func(name string) template.HTML {
		return template.HTML("<b>" + name + "</b>") // trust html
	})

	iris.UseTemplateFunc("system_time", func() string {
		return time.Now().Format(time.RFC822)
	})
	// note that, these template funcs exists for ALL template engines, you are not limited to the iris' default
	// you can iris.UseTemplate amber, pug, django, handlebars too.

	// http://localhost:8080/hi/yourname
	iris.Get("/hi/:name", func(ctx *iris.Context) {
		name := ctx.Param("name")
		ctx.MustRender("hi.html", iris.Map{"Name": name})
	})

	iris.Listen(":8080")
}

// other way to set a a funcmap per template engine and not shared, is from each template engine's configuration,
// all of them has a Funcs field, see other examples in the parent folder.
