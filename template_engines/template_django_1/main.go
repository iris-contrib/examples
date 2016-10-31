package main

import (
	"gopkg.in/kataras/go-template.v0/django"
	"gopkg.in/kataras/iris.v4"
)

func main() {
	iris.UseTemplate(django.New()).Directory("./templates", ".html")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.MustRender("mypage.html", map[string]interface{}{"username": "iris", "is_admin": true}, iris.RenderOptions{"gzip": true})
	})

	iris.Listen(":8080")
}
