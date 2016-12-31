package main

import (
	"github.com/kataras/iris"
)

func main() {

	iris.Get("/set", func(c *iris.Context) {
		c.Session().SetFlash("name", "iris")
		c.Writef("Message setted, is available for the next request")
	})

	iris.Get("/get", func(c *iris.Context) {
		name := c.Session().GetFlashString("name")
		if name != "" {
			c.Writef("Empty name!!")
			return
		}
		c.Writef("Hello %s", name)
	})

	iris.Get("/test", func(c *iris.Context) {
		name := c.Session().GetFlashString("name")
		if name != "" {
			c.Writef("Empty name!!")
			return
		}

		c.Writef("Ok you are comming from /set ,the value of the name is %s", name)
		c.Writef(", and again from the same context: %s", name)

	})

	iris.Listen(":8080")
}
