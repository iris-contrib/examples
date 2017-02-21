package main

import (
	"gopkg.in/kataras/iris.v6"
)

func main() {

	//	usersAPI := iris.None("/api/users/:userid", func(ctx *iris.Context) {
	//		ctx.Writef("user with id: %s", ctx.Param("userid"))
	//	})("api.users.id") // we need to call empty ("") in order to get its iris.Route instance
	//	// or ("the name of the route")
	//	// which later on can be found with iris.Lookup("the name of the route")

	//	static := iris.StaticHandler("/", "./www", false, false)
	//	iris.Get("/*file", iris.Prioritize(usersAPI), static)
	// ___OR___
	// static := iris.NewStaticHandlerBuilder("./www").Path("/").Except(usersAPI).Build()
	// iris.Get("/*file", static)

	// ___OR___SIMPLY:

	usersAPI := iris.None("/api/users/:userid", func(ctx *iris.Context) {
		ctx.Writef("user with id: %s", ctx.Param("userid"))
	})("api.users.id")

	iris.StaticWeb("/", "./www", usersAPI)

	//
	// START THE SERVER
	//
	iris.Listen("localhost:8080")
}
