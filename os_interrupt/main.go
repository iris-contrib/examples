// Run this on command like with: go run main.go , after 2-3 seconds press control+C
package main

import (
	"time"

	"github.com/kataras/iris"
)

func main() {

	iris.Plugins.PostInterrupt(func(s *iris.Framework) {
		// when os.Interrupt signal is fired the body of this function will be
		// fired,
		// you're responsible for closing the server with s.Close()

		// if that event is not registered then the framework
		// will close the server for you.

		/* Do  any custom cleanup and finally call the s.Close()
		   remember you have the iris.Plugins.PreClose(func(s *Framework)) event
		   too
		   so you can split your logic in two logically places.
		*/

		println("control+C pressed, closing the server in 5 seconds!")

		time.Sleep(5 * time.Second)

		s.Close() // ln.Close(), or s.Close(), doesn't matters.
	})

	iris.Get("/", func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "<h1>Hello from index!</h1>")
	})

	// make use of our custom tcp listener here
	// in order to not panic on .Close() (.Listen functions panics on error, .Serve is not, you are responsible)
	ln, err := iris.TCP4(":8080")
	if err != nil {
		panic(err)
	}
	iris.Serve(ln)

}
