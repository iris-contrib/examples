package main

import (
	"time"

	"github.com/iris-contrib/middleware/basicauth"
	"github.com/kataras/iris"
)

func main() {
	authConfig := basicauth.Config{
		Users:      map[string]string{"myusername": "mypassword", "mySecondusername": "mySecondpassword"},
		Realm:      "Authorization Required", // if you don't set it it's "Authorization Required"
		ContextKey: "mycustomkey",            // if you don't set it it's "user"
		Expires:    time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	// to global iris.Use(authentication)
	// to routes
	/*
		iris.Get("/mysecret", authentication, func(ctx *iris.Context) {
			username := ctx.GetString("mycustomkey") //  the Contextkey from the authConfig
			ctx.Write("Hello authenticated user: %s ", username)
		})
	*/

	// to party

	needAuth := iris.Party("/secret", authentication)
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

	iris.Listen(":8080")
}
