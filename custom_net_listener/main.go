package main

import (
	"net"

	"gopkg.in/kataras/iris.v6"
)

func main() {
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello from the server")
	})

	iris.Get("/mypath", func(ctx *iris.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})

	// create our custom listener
	ln, err := net.Listen("tcp4", ":8080")
	if err != nil {
		panic(err)
	}

	// use of the custom listener
	iris.Serve(ln)

	// now if you navigate to http://127.0.0.1:8080/mypath
}
