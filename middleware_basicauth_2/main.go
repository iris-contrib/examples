package main

import (
	"time"

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

	authConfig := basicauth.Config{
		Users:      map[string]string{"myusername": "mypassword", "mySecondusername": "mySecondpassword"},
		Realm:      "Authorization Required", // if you don't set it it's "Authorization Required"
		ContextKey: "mycustomkey",            // if you don't set it it's "user"
		Expires:    time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	// to global app.Use(authentication)
	// to routes
	/*
		app.Get("/mysecret", authentication, func(ctx *iris.Context) {
			username := ctx.GetString("mycustomkey") //  the Contextkey from the authConfig
			ctx.Writef("Hello authenticated user: %s ", username)
		})
	*/

	// to party

	needAuth := app.Party("/secret", authentication)
	{
		needAuth.Get("/", func(ctx *iris.Context) {
			username := ctx.GetString("mycustomkey") //  the Contextkey from the authConfig
			ctx.Writef("Hello authenticated user: %s from localhost:8080/secret ", username)
		})

		needAuth.Get("/profile", func(ctx *iris.Context) {
			username := ctx.GetString("mycustomkey") //  the Contextkey from the authConfig
			ctx.Writef("Hello authenticated user: %s from localhost:8080/secret/profile ", username)
		})

		needAuth.Get("/settings", func(ctx *iris.Context) {
			username := authConfig.User(ctx) // same thing as ctx.GetString("mycustomkey")
			ctx.Writef("Hello authenticated user: %s from localhost:8080/secret/settings ", username)
		})
	}

	app.Listen(":8080")
}
