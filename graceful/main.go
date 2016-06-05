package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/graceful"
)

func main() {
	api := iris.New()
	api.Get("/", func(c *iris.Context) {
		c.Write("Welcome to the home page!")
	})

	graceful.Run(":3001", time.Duration(10)*time.Second, api)
}
