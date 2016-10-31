package main

// History:
// As of, Iris version 4.2.7, you can now specify a custom Handler to create your own router, without the default behavior (subdomains, parameters, autocorrect path)
// This normally is not nessecary but maybe it's useful for some sittuations when you have your main iris server with the default iris router
// and you want to create a new iris instance with a new router which can do other things before serving this request to the default iris instace
//                                                (but in this example we will not cover this usage (you can look multiserver_listening examples))

import (
	"bytes"

	"gopkg.in/kataras/iris.v4"
	"github.com/valyala/fasthttp"
)

var helloPath = []byte("/hello")
var message = []byte("hello world")

func main() {

	api := iris.New()

	api.Router = func(ctx *fasthttp.RequestCtx) {
		if bytes.Equal(ctx.Method(), iris.MethodGetBytes) {
			if bytes.Equal(ctx.Path(), helloPath) {
				ctx.Write(message)
			} else {
				ctx.Error(iris.StatusText(iris.StatusNotFound), iris.StatusNotFound)
			}
		} else {
			ctx.Error(iris.StatusText(iris.StatusMethodNotAllowed), iris.StatusMethodNotAllowed)
		}
	}

	// Go to http://localhost:8080/hello
	api.Listen(":8080")
}
