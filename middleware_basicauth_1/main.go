package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
)

func main() {
	authentication := basicauth.Default(map[string]string{"myusername": "mypassword", "mySecondusername": "mySecondpassword"})

	// to global iris.UseFunc(authentication)
	// to party: iris.Party("/secret", authentication) { ... }

	// to routes
	iris.Get("/mysecret", authentication, func(ctx *iris.Context) {
		username := ctx.GetString("auth") // this can be changed, you will see at the middleware_basic_auth_2 folder
		ctx.Write("Hello authenticated user: %s ", username)
	})

	iris.Listen(":8080")
}
