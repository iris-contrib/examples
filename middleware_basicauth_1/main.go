package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/basicauth"
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())
	authentication := basicauth.Default(map[string]string{"myusername": "mypassword", "mySecondusername": "mySecondpassword"})

	// to global app.Use(authentication)
	// to party: app.Party("/secret", authentication) { ... }

	// to routes
	app.Get("/secret", authentication, func(ctx *iris.Context) {
		username := ctx.GetString("user") // this can be changed, you will see at the middleware_basic_auth_2 folder
		ctx.Writef("Hello authenticated user: %s ", username)
	})

	app.Get("/secret/profile", authentication, func(ctx *iris.Context) {
		username := ctx.GetString("user")
		ctx.Writef("Hello authenticated user: %s from localhost:8080/secret/profile ", username)
	})

	app.Get("/othersecret", authentication, func(ctx *iris.Context) {
		username := ctx.GetString("user")
		ctx.Writef("Hello authenticated user: %s from localhost:8080/othersecret ", username)
	})

	app.Listen(":8080")
}
