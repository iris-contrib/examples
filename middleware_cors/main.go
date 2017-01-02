package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Use(cors.Default())

	app.Post("/user/create", func(ctx *iris.Context) {
		ctx.Writef("POST CREATE")
	}) //db.insert...}
	app.Put("/user/edit/:email", func(ctx *iris.Context) {
		ctx.Writef("PUT EDIT EMAIL: %s", ctx.Param("email"))
	}) //db.update...}

	app.Listen(":3333")
	//	or
	//	ln, err := iris.TCP4("localhost:3333")
	//	if err != nil {
	//		panic(err)
	//	}

	//	app.Serve(ln)
}
