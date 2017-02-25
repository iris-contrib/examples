package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/recover"
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	app.Use(recover.New())
	i := 0
	app.Get("/", func(ctx *iris.Context) {
		i++
		if i%2 == 0 {
			panic("a panic here")
		}

		ctx.Next()

	}, func(ctx *iris.Context) {
		ctx.Writef("Hello, refresh one time more to get panic!")
	})

	// open http://localhost:8080, and hit refresh, each two refreshes you'll get a panic follows by recovery
	app.Listen(":8080")
}
