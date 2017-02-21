package main

import (
	"github.com/iris-contrib/middleware/cors"
	"gopkg.in/kataras/iris.v6"
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
}
