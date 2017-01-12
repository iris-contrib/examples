package main

import (
	"github.com/kataras/iris"
)

func main() {

	iris.None("/api/user/:userid", func(ctx *iris.Context) {
		userid := ctx.Param("userid")
		ctx.Writef("user with id: %s", userid)
	})("user.api")

	// change the "user.api" state from offline to online and online to offline
	iris.Get("/change", func(ctx *iris.Context) {
		routeName := "user.api"
		if iris.Lookup(routeName).IsOnline() {
			// set to offline
			iris.SetRouteOffline(routeName)
		} else {
			// set to online if it was not online(so it was offline)
			iris.SetRouteOnline(routeName, iris.MethodGet)
		}
	})

	//	iris.Get("/execute/:routename", func(ctx *iris.Context) {
	//		routeName := ctx.Param("routename")
	//		userAPICtx := ctx.ExecuteRoute(routeName)
	//		if userAPICtx == nil {
	//			ctx.Writef("Route with name: %s didnt' found or couldn't be validate with this request path!", routeName)
	//		}
	//	})

	iris.Get("/execute", func(ctx *iris.Context) {
		routeName := "user.api"
		// change the path in order to be catcable from the ExecuteRoute
		// ctx.Request.URL.Path = "/api/user/42"
		// ctx.ExecRoute(routeName)
		// or:
		ctx.ExecRouteAgainst(routeName, "/api/user/42")
	})

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello from index /")
	})

	//
	// START THE SERVER
	//
	// STEPS:
	// 1. navigate to http://localhost:8080/user/api/42
	// you should get 404 error
	// 2. now, navigate to http://localhost:8080/change
	// you should see a blank page
	// 3. now, navigate to http://localhost:8080/user/api/42
	// you should see the page working, NO 404 error
	// go back to the http://localhost:8080/change
	// you should get 404 error again
	// You just dynamically changed the state of a route with 3 lines of code!
	// you can do the same with group of routes and subdomains :)
	iris.Listen(":8080")
}
