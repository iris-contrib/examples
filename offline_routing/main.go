package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	// You can set a name to a route by app.Get("/mypath", handler).ChangeName("theRouteName")
	// You can get a RouteInfo by app.Routes().Lookup("theRouteName")
	userAPI := app.None("/api/users/:userid", func(ctx *iris.Context) {
		userid := ctx.Param("userid")
		ctx.Writef("user with id: %s", userid)
	})

	// change the "users.api" state from offline to online and online to offline
	app.Get("/change", func(ctx *iris.Context) {
		if userAPI.IsOnline() {
			app.Routes().Offline(userAPI)
		} else {
			// set to online if it was not online(so it was offline)
			app.Routes().Online(userAPI, iris.MethodGet)
		}
	})

	app.Get("/execute", func(ctx *iris.Context) {
		// change the path in order to be catcable from the ExecuteRoute
		// ctx.Request.URL.Path = "/api/users/42"
		// ctx.ExecRoute(iris.Route)
		// or:
		ctx.ExecRouteAgainst(userAPI, "/api/users/42")
	})

	app.Get("/", func(ctx *iris.Context) {
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
	// Go to http://localhost:8080/execute to execute the offline route and send its result!
	//
	// You just dynamically changed the state of a route with 3 lines of code!
	// you can do the same with group of routes and subdomains, have fun :)
	app.Listen(":8080")
}
