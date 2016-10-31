package main

import "gopkg.in/kataras/iris.v4"

func main() {

	iris.Get("/servezip", func(c *iris.Context) {
		file := "./files/first.zip"
		c.SendFile(file, "first.zip")
	})

	iris.Listen(":8080")
}
