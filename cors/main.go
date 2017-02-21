package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	app.Adapt(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	app.Post("/user/create", func(ctx *iris.Context) {
		ctx.Writef("POST CREATE")
	}) //db.insert...}
	app.Put("/user/edit/:email", func(ctx *iris.Context) {
		ctx.Writef("PUT EDIT EMAIL: %s", ctx.Param("email"))
	}) //db.update...}

	app.Listen(":3333")
}
