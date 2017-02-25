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

	// Match to /hello/iris,  (if PathCorrection:true match also /hello/iris/)
	// Not match to /hello or /hello/ or /hello/iris/something
	app.Get("/hello/:name", func(ctx *iris.Context) {
		// Retrieve the parameter name
		name := ctx.Param("name")
		ctx.Writef("Hello %s", name)
	})

	// Match to /profile/iris/friends/1, (if PathCorrection:true match also /profile/iris/friends/1/)
	// Not match to /profile/ , /profile/iris ,
	// Not match to /profile/iris/friends,  /profile/iris/friends ,
	// Not match to /profile/iris/friends/2/something
	app.Get("/profile/:fullname/friends/:friendID", func(ctx *iris.Context) {
		// Retrieve the parameters fullname and friendID
		fullname := ctx.Param("fullname")
		friendID, _ := ctx.ParamInt("friendID")

		ctx.Writef("hello %s with :friendID = %d", fullname, friendID)
	})

	/* Example: /posts/:id and /posts/new (dynamic value conficts with the static 'new') for performance reasons and simplicity
	   but if you need to have them you can do that: */

	app.Get("/posts/*action", func(ctx *iris.Context) {
		action := ctx.Param("action")
		if action == "/new" {
			// it's posts/new page
			ctx.Writef("POSTS NEW")
		} else {
			ctx.Writef("OTHER POSTS")
			// it's posts/:id page
			//doSomething with the action which is the id
		}
	})

	app.Listen(":8080")
}
