package main

import "github.com/kataras/iris"

func main() {
	message := []byte("hello world")
	iris.Get("/hello", func(ctx *iris.Context) {
		ctx.SetBody(message)
	})

	iris.Listen(":8080")
}
