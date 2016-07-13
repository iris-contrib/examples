package main

import (
	"github.com/iris-contrib/template/amber"
	"github.com/kataras/iris"
)

type mypage struct {
	Name string
}

func main() {

	iris.UseEngine(amber.New()).Directory("./templates", ".amber")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("basic.amber", mypage{"iris"}, iris.RenderOptions{"gzip": true})
	})

	iris.Listen(":8080")
}
