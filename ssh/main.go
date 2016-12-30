// Basic remote control of your Iris station via SSH
package main

import "gopkg.in/kataras/iris.v5"

func main() {
	iris.Config.IsDevelopment = true // some logs for ssh if you enabled it.

	// don't try something like this:
	//	iris.SSH = &iris.SSHServer{Host: "0.0.0.0:22", KeyPath: "./iris_rsa", Users: iris.Users{"kataras": []byte("pass")}} //Bin: "C:/Program Files/Git/usr/bin"
	// instead:
	iris.SSH.Host = "0.0.0.0:22"
	iris.SSH.KeyPath = "./iris_rsa" // it's auto-generated if not exists
	iris.SSH.Users = iris.Users{"kataras": []byte("pass")}

	// if you want the ssh logger to catch the errors you have to register them, even with empty handler
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusNotFound, "<b> NOT FOUND </b>")
	})

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello from HTTP Server, restart your browser's tab if you 'stop' the server but this stills to be shown [browser caches the connection].")
	})

	iris.Listen(":8080")
}
