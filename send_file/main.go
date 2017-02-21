package main

import "gopkg.in/kataras/iris.v6"

func main() {

	iris.Get("/servezip", func(c *iris.Context) {
		file := "./files/first.zip"
		c.SendFile(file, "c.zip")
	})

	iris.Listen(":8080")
}
