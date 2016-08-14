// For automatic TLS look at /letsencrypt folder
package main

import (
	"github.com/kataras/iris"
)

/*
	Previous example: multiserver_listening
*/
func main() {
	host := "127.0.0.1:443"
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello from the SECURE server")
	})

	iris.Get("/mypath", func(ctx *iris.Context) {
		ctx.Write("Hello from the SECURE server on path /mypath")
	})

	iris.Get("/redirect", func(ctx *iris.Context) {
		ctx.Redirect("/home")
	})

	iris.Get("/home", func(ctx *iris.Context) {
		ctx.Write("Hello from %s", ctx.PathString())
	})

	// start the MAIN server (HTTPS) on port 443, this is a blocking func
	iris.ListenTLS(host, "mycert.cert", "mykey.key")
}

// For automatic TLS look at /letsencrypt folder
