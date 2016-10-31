package main

import (
	"gopkg.in/iris-contrib/middleware.v4/cors"
	"gopkg.in/kataras/iris.v4"
)

func main() {
	app := iris.New()
	app.Use(cors.Default())

	app.Post("/user/create", func(ctx *iris.Context) {
		ctx.Write("POST CREATE")
	}) //db.insert...}
	app.Put("/user/edit/:email", func(ctx *iris.Context) {
		ctx.Write("PUT EDIT EMAIL: %s", ctx.Param("email"))
	}) //db.update...}

	app.Listen(":3333")
	//	or
	//	ln, err := iris.TCP4("localhost:3333")
	//	if err != nil {
	//		panic(err)
	//	}

	//	app.Serve(ln)
}
