package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	server1 := iris.New()
	server1.Adapt(httprouter.New())
	server1.Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello from the server1 on :8080")
	})

	server2 := iris.New()
	server2.Adapt(httprouter.New())
	server2.Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello from the server2 on :80")
	})

	// remember that .Listen on Iris is a block function
	//, so it's blocking, you have to run it in go routine
	// when ever you want your code below that to be executed
	go server1.Listen(":8080")
	server2.Listen(":1993")

}
