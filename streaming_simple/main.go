package main

import (
	"fmt" // just an optional helper
	"io"
	"time" // showcase the delay

	"gopkg.in/kataras/iris.v6"
)

func main() {
	timeWaitForCloseStream := 4 * time.Second

	iris.Get("/", func(ctx *iris.Context) {
		i := 0
		// goroutine in order to no block and just wait,
		// goroutine is OPTIONAL and not a very good option but it depends on the needs
		// Look the streaming_simple_2 for an alternative code style
		// Send the response in chunks and wait for a second between each chunk.
		go ctx.StreamWriter(func(w io.Writer) bool {
			i++
			fmt.Fprintf(w, "this is a message number %d\n", i) // write
			time.Sleep(time.Second)                            // imaginary delay

			return true // continue and flush
		})

		// when this handler finished the client should be see the stream writer's contents
		// simulate a job here...
		time.Sleep(timeWaitForCloseStream)
	})

	// or iris.ListenLETSENCRYPT to (logically) enable http/2 flush on stream writer
	iris.Listen(":8080")
}
