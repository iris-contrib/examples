package main

import (
	"github.com/kataras/iris"
)

func main() {

	iris.Get("/set", func(c *iris.Context) {
		c.SetFlash("name", "iris")
	})

	iris.Get("/get", func(c *iris.Context) {
		c.Write("Hello %s", c.GetFlash("name"))
	})

	iris.Get("/test", func(c *iris.Context) {

		name := c.GetFlash("name")
		if name == "" {
			c.Write("Ok you are comming from /get")
		} else {
			c.Write("Ok you are comming from /set ,the value of the name is %s", name)
			c.Write(", and again from the same context: %s", c.GetFlash("name"))
		}
	})

	iris.Listen(":8080")
}
