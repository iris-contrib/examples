package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
)

/*
NOTES:
This is useful only when you need something like that: https://github.com/kataras/iris/issues/235

the main server should always defined last.
*/
func main() {
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello from the server")
	})

	// start a secondary server (HTTP) on port 9090, this is a non-blocking func
	iris.SecondaryListen(config.Server{ListeningAddr: ":9090"})

	// start a secondary server (HTTPS) on port 443, this is a non-blocking func
	iris.SecondaryListen(config.Server{ListeningAddr: ":443", CertFile: "mycert.cert", KeyFile: "mykey.key"}) // you can close this server with .Close()

	// start the MAIN server (HTTP) on port 8080, this is a blocking func
	iris.Listen(":8080")

	// go to the second example 'multiserver_listening2' if you want to see how you can easly redirect from http to https with a second serve

}
