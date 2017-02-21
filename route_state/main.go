package main

import (
	"gopkg.in/kataras/iris.v6"
)

func main() {

	// You can find the Route by iris.Lookup("theRouteName")
	// you can set a route name as: myRoute := iris.Get("/mypath", handler)("theRouteName")
	// that will set a name to the route and returns its iris.Route instance for further usage.
	api := iris.None("/api/users/:userid", func(ctx *iris.Context) {
		userid := ctx.Param("userid")
		ctx.Writef("user with id: %s", userid)
	})("users.api")

	// change the "users.api" state from offline to online and online to offline
	iris.Get("/change", func(ctx *iris.Context) {
		if api.IsOnline() {
			// set to offline
			iris.SetRouteOffline(api)
		} else {
			// set to online if it was not online(so it was offline)
			iris.SetRouteOnline(api, iris.MethodGet)
		}
	})

	iris.Get("/execute", func(ctx *iris.Context) {
		// change the path in order to be catcable from the ExecuteRoute
		// ctx.Request.URL.Path = "/api/users/42"
		// ctx.ExecRoute(iris.Route)
		// or:
		ctx.ExecRouteAgainst(api, "/api/users/42")
	})

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello from index /")
	})

	//
	// START THE SERVER
	//
	// STEPS:
	// 1. navigate to http://localhost:8080/api/users/42
	// you should get 404 error
	// 2. now, navigate to http://localhost:8080/change
	// you should see a blank page
	// 3. now, navigate to http://localhost:8080/api/users/42
	// you should see the page working, NO 404 error
	// go back to the http://localhost:8080/change
	// you should get 404 error again
	// You just dynamically changed the state of a route with 3 lines of code!
	// you can do the same with group of routes and subdomains :)
	iris.Listen(":8080")
}
