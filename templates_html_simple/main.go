package main

import "github.com/kataras/iris"

func main() {
	iris.Get("/hi", hi)
	iris.Listen(":8080")
}

func hi(ctx *iris.Context) {
	ctx.Render("hi.html", struct{ Name string }{Name: "iris"})
}
