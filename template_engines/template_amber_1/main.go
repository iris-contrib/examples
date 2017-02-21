package main

import (
	"github.com/kataras/go-template/amber"
	"gopkg.in/kataras/iris.v6"
)

type mypage struct {
	Name string
}

func main() {

	iris.UseTemplate(amber.New()).Directory("./templates", ".amber")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("basic.amber", mypage{"iris"}, iris.RenderOptions{"gzip": true})
	})

	iris.Listen(":8080")
}
