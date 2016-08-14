package main

import "github.com/kataras/iris"

func main() {

	iris.Get("/servezip", func(c *iris.Context) {
		file := "./files/first.zip"
		// to support resume use the c.RequestCtx.SendFile
		err := c.SendFile(file, "first.zip")
		if err != nil {
			println("error: " + err.Error())
		}
	})

	iris.Listen(":8080")
}
