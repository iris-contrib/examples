package main

import (
	"fmt"

	"github.com/kataras/iris"
)

/* Native messages no need to import the iris-ws.js to the ./templates.client.html
Use of: OnMessage and EmitMessage
*/

type clientPage struct {
	Title string
	Host  string
}

func main() {

	iris.StaticWeb("/js", "./static/js")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.Host()})
	})

	// the path which the websocket client should listen/registed to ->
	iris.Config.Websocket.Endpoint = "/my_endpoint"

	// IF you work with something like proto, enable BinaryMessages:
	// iris.Config.Websocket.BinaryMessages = true

	ws := iris.Websocket // get the websocket server

	ws.OnConnection(func(c iris.WebsocketConnection) {

		c.OnMessage(func(data []byte) {
			message := string(data)
			c.To(iris.Broadcast).EmitMessage([]byte("Message from: " + c.ID() + "-> " + message)) // broadcast to all clients except this
			c.EmitMessage([]byte("Me: " + message))                                               // writes to itself
		})

		c.OnDisconnect(func() {
			fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
		})
	})

	iris.Listen(":8080")
}
