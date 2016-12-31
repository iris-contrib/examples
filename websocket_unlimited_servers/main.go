package main

import (
	"fmt"

	"github.com/kataras/iris"
)

type clientPage struct {
	Title string
	Host  string
}

func main() {
	api := iris.New()

	api.Static("/js", "./static/js", 1)

	api.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.Host()})
	})

	// important staff
	ws := iris.NewWebsocketServer(api)
	// the path which the websocket client should listen/registed to
	ws.Config = iris.WebsocketConfiguration{Endpoint: "/my_endpoint"}

	// more than one custom websocket  server in the same iris instance:
	// entirely new websocket server with its own connections:
	// ws2 := iris.NewWebsocketServer(api)
	// ws2.Config = iris.WebsocketConfiguration{Endpoint: "/my_second_endpoint"}

	var myChatRoom = "room1"
	ws.OnConnection(func(c iris.WebsocketConnection) {

		c.Join(myChatRoom)

		c.On("chat", func(message string) {
			// to all except this connection ->
			//c.To(websocket.Broadcast).Emit("chat", "Message from: "+c.ID()+"-> "+message)

			// to the client ->
			//c.Emit("chat", "Message from myself: "+message)

			//send the message to the whole room,
			//all connections are inside this room will receive this message
			c.To(myChatRoom).Emit("chat", "From: "+c.ID()+": "+message)
		})

		c.OnDisconnect(func() {
			fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
		})
	})

	//

	api.Listen(":8080")
}
