package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recovery"
)

func main() {
	//iris.Use(recovery.Recovery(os.Stdout)) // this is an optional parameter, you can skip it, the default is os.Stderr
	iris.Use(recovery.Recovery())
	i := 0
	iris.Get("/", func(ctx *iris.Context) {
		i++
		if i%2 == 0 {
			panic("a panic here")
			return

		}

		ctx.Next()

	}, func(ctx *iris.Context) {
		ctx.Write("Hello, refresh one time more to get panic!")
	})

	iris.Listen(":8080")
}
