package main

import (
	"fmt"
	"time"

	"github.com/kataras/iris"
)

func main() {

	iris.UseFunc(responseLogger) // global middleware, catch all routes

	iris.Get("/", func(c *iris.Context) {
		c.HTML(iris.StatusOK, "<h1> Hello from / </h1>")
	})

	iris.Get("/home", func(c *iris.Context) {
		c.HTML(iris.StatusOK, "<h1> Hello from /home</h1>")
	})

	iris.Listen(":8080")
}

func responseLogger(c *iris.Context) {
	c.Next() // process the request first, we don't want to have delays

	date := time.Now().Format("01/02 - 15:04:05")
	fmt.Printf("%s\n%s", date, c.Response.String())
}
