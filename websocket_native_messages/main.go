package main

import (
	"fmt"

	"gopkg.in/kataras/iris.v5"
)

/* Native messages no need to import the iris-ws.js to the ./templates.client.html
Use of: OnMessage and EmitMessage
*/

type clientPage struct {
	Title string
	Host  string
}

func main() {

	iris.Static("/js", "./static/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.HostString()})
	})

	// the path which the websocket client should listen/registed to ->
	iris.Config.Websocket.Endpoint = "/my_endpoint"

	// IF you work with something like proto, enable BinaryMessages:
	// iris.Config.Websocket.BinaryMessages = true

	ws := iris.Websocket // get the websocket server

	ws.OnConnection(func(c iris.WebsocketConnection) {

		c.OnMessage(func(data []byte) {
			message := string(data)
			c.To(iris.Broadcast).EmitMessage([]byte("Message from: " + c.ID() + "-> " + message))
			c.EmitMessage([]byte("Me: " + message))
		})

		c.OnDisconnect(func() {
			fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
		})
	})

	iris.Listen(":8080")
}
