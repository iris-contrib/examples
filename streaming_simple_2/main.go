package main

import (
	"fmt" // just an optional helper
	"io"
	"time" // showcase the delay

	"gopkg.in/kataras/iris.v6"
)

func main() {

	iris.Get("/", func(ctx *iris.Context) {

		// Send the response in chunks and wait for a second between each chunk.
		ctx.StreamWriter(func(w io.Writer) bool {
			for i := 0; i <= 4; i++ {
				fmt.Fprintf(w, "this is a message number %d\n", i) // write
				time.Sleep(time.Second)
			}

			// when this handler finished the client should be see the stream writer's contents
			return false // stop and flush the contents
		})

	})

	// or iris.ListenLETSENCRYPT to (logically) enable http/2 flush on stream writer
	iris.Listen(":8080")
}
