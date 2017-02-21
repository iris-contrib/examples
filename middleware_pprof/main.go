package main

import (
	"github.com/iris-contrib/middleware/pprof"
	"gopkg.in/kataras/iris.v6"
)

func main() {
	iris.Get("/", func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "<h1> Please click <a href='/debug/pprof'>here</a>")
	})

	iris.Get("/debug/pprof/*action", pprof.New())

	iris.Listen(":8080")
}
