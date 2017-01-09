package main

import (
	"fmt"     // optional
	"strings" // optional

	"github.com/kataras/iris"
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

	var myChatRoom = "room1"
	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
		// Request returns the (upgraded) *http.Request of this connection
		// avoid using it, you normally don't need it,
		// websocket has everything you need to authenticate the user BUT if it's necessary
		// then  you use it to receive user information, for example: from headers.

		httpRequest := c.Request()
		fmt.Printf("Headers for the connection with ID: %s\n\n", c.ID())
		for k, v := range httpRequest.Header {
			fmt.Printf("%s = '%s'\n", k, strings.Join(v, ", "))
		}
		c.Join(myChatRoom)

		c.On("chat", func(message string) {
			if message == "leave" {
				c.Leave(myChatRoom)
				c.To(myChatRoom).Emit("chat", "Client with ID: "+c.ID()+" left from the room and cannot send or receive message to/from this room.")
				c.Emit("chat", "You have left from the room: "+myChatRoom+" you cannot send or receive any messages from others inside that room.")
				return
			}
			// to all except this connection ->
			// c.To(iris.Broadcast).Emit("chat", "Message from: "+c.ID()+"-> "+message)
			// to all connected clients: c.To(iris.All)

			// to the client itself ->
			//c.Emit("chat", "Message from myself: "+message)

			//send the message to the whole room,
			//all connections are inside this room will receive this message
			c.To(myChatRoom).Emit("chat", "From: "+c.ID()+": "+message)
		})

		// or create a new leave event
		// c.On("leave", func() {
		// 	c.Leave(myChatRoom)
		// })

		c.OnDisconnect(func() {
			fmt.Printf("Connection with ID: %s has been disconnected!\n", c.ID())
		})
	})

	iris.Listen(":8080")
}
