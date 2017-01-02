package main

import (
	"time"

	"github.com/iris-contrib/graceful"
	"github.com/kataras/iris"
)

func main() {
	api := iris.New()
	api.Get("/", func(c *iris.Context) {
		c.Writef("Welcome to the home page!")
	})

	graceful.Run(":3001", time.Duration(20)*time.Second, api)
}
