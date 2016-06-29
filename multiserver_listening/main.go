package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
)

/*
NOTES:
This is useful only when you need that
all Listen functions are blocking the execution so run all in goroutine except the last server instance you want to listen to

*/
func main() {
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello from the server")
	})

	// start the first server (HTTP) on port 8080
	go iris.Listen(":8080")
	// start a second server (HTTP) on port 9090
	go iris.ListenToServer(config.Server{ListeningAddr: ":9090"})

	// start a third server (HTTPS) on port 443
	mylastServer, err := iris.ListenToServer(config.Server{ListeningAddr: ":443", CertFile: "mycert.cert", KeyFile: "mykey.key"})
	// you can close this server with mylastServer.Close()
}
