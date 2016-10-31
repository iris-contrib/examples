package main

import (
	"gopkg.in/kataras/iris.v4"
)

func main() {

	iris.Get("/set", func(c *iris.Context) {
		c.SetFlash("name", "iris")
		c.Write("Message setted, is available for the next request")
	})

	iris.Get("/get", func(c *iris.Context) {
		name, err := c.GetFlash("name")
		if err != nil {
			c.Write(err.Error())
			return
		}
		c.Write("Hello %s", name)
	})

	iris.Get("/test", func(c *iris.Context) {

		name, err := c.GetFlash("name")
		if err != nil {
			c.Write(err.Error())
			return
		}

		c.Write("Ok you are comming from /set ,the value of the name is %s", name)
		c.Write(", and again from the same context: %s", name)

	})

	iris.Listen(":8080")
}
