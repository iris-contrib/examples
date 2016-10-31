package main

import (
	"strconv"

	"gopkg.in/kataras/iris.v4"
)

func main() {

	// Match to /hello/iris,  (if PathCorrection:true match also /hello/iris/)
	// Not match to /hello or /hello/ or /hello/iris/something
	iris.Get("/hello/:name", func(c *iris.Context) {
		// Retrieve the parameter name
		name := c.Param("name")
		c.Write("Hello %s", name)
	})

	// Match to /profile/iris/friends/1, (if PathCorrection:true match also /profile/iris/friends/1/)
	// Not match to /profile/ , /profile/iris ,
	// Not match to /profile/iris/friends,  /profile/iris/friends ,
	// Not match to /profile/iris/friends/2/something
	iris.Get("/profile/:fullname/friends/:friendID", func(c *iris.Context) {
		// Retrieve the parameters fullname and friendID
		fullname := c.Param("fullname")
		friendID, err := c.ParamInt("friendID")
		if err != nil {
			// Do something with the error
		}
		c.HTML(iris.StatusOK, "<b> Hello </b>"+fullname+"<b> with friends ID </b>"+strconv.Itoa(friendID))
	})

	/* Example: /posts/:id and /posts/new (dynamic value conficts with the static 'new') for performance reasons and simplicity
	   but if you need to have them you can do that: */

	iris.Get("/posts/*action", func(ctx *iris.Context) {
		action := ctx.Param("action")
		if action == "/new" {
			// it's posts/new page
			ctx.Write("POSTS NEW")
		} else {
			ctx.Write("OTHER POSTS")
			// it's posts/:id page
			//doSomething with the action which is the id
		}
	})

	iris.Listen(":8080")
}
