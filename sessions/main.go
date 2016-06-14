package main

import (
	"github.com/kataras/iris"
)

// default memory provider
func main() {

	iris.Get("/set", func(c *iris.Context) {

		//set session values
		c.Session().Set("name", "iris")

		//test if setted here
		c.Write("All ok session setted to: %s", c.Session().GetString("name"))
	})

	iris.Get("/get", func(c *iris.Context) {
		name := c.Session().GetString("name")

		c.Write("The name on the /set was: %s", name)
	})

	iris.Get("/delete", func(c *iris.Context) {
		//get the session for this context

		c.Session().Delete("name")

	})

	iris.Get("/clear", func(c *iris.Context) {

		// removes all entries
		c.Session().Clear()
	})

	iris.Get("/destroy", func(c *iris.Context) {
		//destroy, removes the entire session and cookie
		c.SessionDestroy()
	})

	iris.Listen(":8080")
}
