package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/kataras/iris"
)

type clientPage struct {
	Title string
	Host  string
}

func main() {
	iris.StaticWeb("/js", "./static/js")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.ServerHost()})
	})

	iris.Config.Websocket.Endpoint = "/my_endpoint"
	Conn := make(map[iris.WebsocketConnection]bool)
	var myChatRoom = "room1"
	var mutex = new(sync.Mutex)

	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
		c.Join(myChatRoom)
		mutex.Lock()
		Conn[c] = true
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
			delete(Conn, c)
			mutex.Unlock()
			fmt.Printf("\nConnection with ID: %s has been disconnected!\n", c.ID())
		})
	})

	var delay = 1 * time.Second
	go func() {
		i := 0
		for {
			mutex.Lock()
			broadcast(Conn, fmt.Sprintf("aaaa %d\n", i))
			mutex.Unlock()
			time.Sleep(delay)
			i++
		}
	}()

	go func() {
		i := 0
		for {
			mutex.Lock()
			broadcast(Conn, fmt.Sprintf("aaaa2 %d\n", i))
			mutex.Unlock()
			time.Sleep(delay)
			i++
		}
	}()

	iris.Listen(":8080")
}

func broadcast(Conn map[iris.WebsocketConnection]bool, message string) {
	for k, _ := range Conn {
		k.To("room1").Emit("chat", message)
	}
}
