package main

import (
	"fmt"
	"sync"
	"time"

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
		ctx.Render("client.html", clientPage{"Client Page", ctx.ServerHost()})
	})

	var delay = 1 * time.Second
	go func() {
		i := 0
		for {
			broadcast(fmt.Sprintf("aaaa %d\n", i))
			time.Sleep(delay)
			i++
		}
	}()

	go func() {
		i := 0
		for {
			broadcast(fmt.Sprintf("aaaa2 %d\n", i))
			time.Sleep(delay)
			i++
		}
	}()

	app.Listen(":8080")
}

var myChatRoom = "room1"
var mutex = new(sync.Mutex)
var conn = make(map[websocket.Connection]bool)

func handleWebsocket(c websocket.Connection) {
	c.Join(myChatRoom)
	mutex.Lock()
	conn[c] = true
	mutex.Unlock()
	c.On("chat", func(message string) {
		if message == "leave" {
			c.Leave(myChatRoom)
			c.To(myChatRoom).Emit("chat", "Client with ID: "+c.ID()+" left from the room and cannot send or receive message to/from this room.")
			c.Emit("chat", "You have left from the room: "+myChatRoom+" you cannot send or receive any messages from others inside that room.")
			return
		}
	})
	c.OnDisconnect(func() {
		mutex.Lock()
		delete(conn, c)
		mutex.Unlock()
		fmt.Printf("\nConnection with ID: %s has been disconnected!\n", c.ID())
	})
}
func broadcast(message string) {
	mutex.Lock()
	for k := range conn {
		k.To("room1").Emit("chat", message)
	}
	mutex.Unlock()
}
