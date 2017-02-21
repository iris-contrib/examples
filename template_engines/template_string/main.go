package main

import (
	"github.com/kataras/go-template/django"
	"gopkg.in/kataras/iris.v6"
)

func main() {

	// iris.TemplateString
	// Executes and parses the template but instead of rendering to the client, it returns the contents.
	// Useful when you want to send a template via e-mail or anything you can imagine.

	// Note that: iris.TemplateString can be called outside of the context also

	iris.UseTemplate(django.New()).Directory("./templates", ".html")

	iris.Get("/", func(ctx *iris.Context) {
		// the same you can do with serializers using the iris.SerializeToString)

		rawHtmlContents := iris.TemplateString("mypage.html", map[string]interface{}{"username": "iris", "is_admin": true}, iris.RenderOptions{"charset": "UTF-8"}) // defaults to UTF-8 already
		ctx.Log(rawHtmlContents)
		ctx.Writef("The Raw HTML is:\n%s", rawHtmlContents)
	})

	iris.Listen(":8080")
}
