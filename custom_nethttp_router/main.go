package main

import (
	"github.com/kataras/iris"
	"net/http"
)

func main() {

	api := iris.New()

	api.Router = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// optional: get the context to get its helper functions
		ctx := api.AcquireCtx(w, r)
		if ctx.Path() == "/hello" {
			ctx.WriteString("hello world")
		} else {
			ctx.EmitError(http.StatusNotFound)
		}

		// release the ctx
		api.ReleaseCtx(ctx)
	})

	// Go to http://localhost:8080/hello
	api.Listen(":8080")
}
