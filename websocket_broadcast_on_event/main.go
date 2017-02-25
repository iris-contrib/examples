package main

import (
	"fmt"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"gopkg.in/kataras/iris.v6/adaptors/websocket"
)

type clientPage struct {
	Title string
	Host  string
}

func main() {
	app := iris.New()

	app.Adapt(iris.DevLogger())

	app.Adapt(httprouter.New())

	app.Adapt(view.HTML("./templates", ".html"))
	ws := websocket.New(websocket.Config{
		// the path which the websocket client should listen/registed to ->
		Endpoint: "my_endpoint",
		// WriteTimeout = 60 * time.Second,
		// CheckOrigin: (...)bool, by default all origins are allowed.
	})

	ws.OnConnection(handleWebsocket)

	// adapt the websocket server and you're ready
	app.Adapt(ws)

	app.StaticWeb("/js", "./static/js")

	app.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.Host()})
	})

	app.Listen(":8080")
}

func handleWebsocket(c websocket.Connection) {
	// when event "chat" fired from client side, catch this message and:
	// send it back with the "chat" event (client(.js) also waits for this event) to the client as Me: $themessage$
	// and to all clients except itself as From $thisconnectionid: $themessage
	c.On("chat", func(message string) {
		// to all except this client ->
		c.To(websocket.Broadcast).Emit("chat", "From "+c.ID()+": "+message)

		// to this client as Me:
		c.Emit("chat", "Me :"+message)
	})

	c.OnDisconnect(func() {
		fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
	})
}
