package main

import (
	"fmt"

	"gopkg.in/kataras/iris.v6"
)

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
	// by-default all origins are accepted, you can change this behavior by setting:
	// iris.Config.Websocket.CheckOrigin
	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {

		// when event "chat" fired from client side, catch this message and:
		// send it back with the "chat" event (client(.js) also waits for this event) to the client as Me: $themessage$
		// and to all clients except itself as From $thisconnectionid: $themssage
		c.On("chat", func(message string) {
			// to all except this client ->
			c.To(iris.Broadcast).Emit("chat", "From "+c.ID()+": "+message)

			// to this client as Me:
			c.Emit("chat", "Me :"+message)
		})

		c.OnDisconnect(func() {
			fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
		})
	})

	iris.Listen(":8080")
}
