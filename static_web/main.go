package main

import "github.com/kataras/iris"

func main() {
	//This iris.StaticWeb("/static", "./static", 1)
	// or
	iris.StaticWeb("/", "./static", 0)
	iris.Listen(":8080")
}
